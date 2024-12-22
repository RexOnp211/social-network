package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"social-network/pkg/helpers"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type ClientList map[*Client]bool

type EventHandler func(event Event, c *Client) error

type Manager struct {
	clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
}

type Client struct {
	connection  *websocket.Conn
	manager     *Manager
	egress      chan Event
	user        *helpers.User
	recipientId string
}

func NewClient(conn *websocket.Conn, manager *Manager, userId int, username string) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan Event),
		user: &helpers.User{
			Nickname: username,
			Id:       userId,
		},
	}
}

func (c *Client) Read() {
	defer func() {
		c.manager.RemoveClient(c)
	}()

	c.connection.SetReadLimit(1024)

	if err := c.connection.SetReadDeadline(time.Now().Add(1000 * time.Second)); err != nil {
		log.Printf("Failed to set read deadline: %v", err)
		return
	}

	for {
		fmt.Println("reading message")
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}

		fmt.Println("received message", string(payload))

		var event Event

		if err := json.Unmarshal(payload, &event); err != nil {
			log.Printf("Error unmarshalling event: %v", err)
			return
		}

		if err := c.manager.routeEvent(event, c); err != nil {
			log.Printf("Error routing event: %v", err)
			return
		}
	}
}

func (c *Client) Write() {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer func() {
		ticker.Stop()
		c.connection.Close()
	}()

	for {
		select {
		case event := <-c.egress:
			if err := c.connection.WriteJSON(event); err != nil {
				log.Printf("Error writing JSON event: %v", err)
				return
			}
			data, err := json.Marshal(event)
			if err != nil {
				log.Printf("Error marshalling event: %v", err)
				return
			}
			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Printf("Error writing message: %v", err)
			}
			fmt.Println("sent message")
		case <-ticker.C:
			if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Error sending ping message: %v", err)
				return
			}
		}
	}
}

func NewManager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()
	return m
}

var (
	SendFollowRequest     = "follow_request"
	FollowRequestList     = "follow_request_list"
	Follow_request_status = "follow_request_status"
	RemoveFollowRequest   = "remove_follow_request"
)

func (m *Manager) setupEventHandlers() {
	m.handlers[SendFollowRequest] = SendFollowRequestHandler
	m.handlers[FollowRequestList] = GetFollowRequestsHandler
	m.handlers[Follow_request_status] = AcceptOrDeclineFollowRequest
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return fmt.Errorf("handler error for event type %s: %w", event.Type, err)
		}
	}
	return nil
}

func (m *Manager) AddClient(c *Client) {
	log.Println("adding client to manager:", c.user.Nickname)
	m.Lock()
	defer m.Unlock()
	m.clients[c] = true
}

func (m *Manager) RemoveClient(c *Client) {
	log.Println("removing client from manager:", c.user.Nickname)
	if _, ok := m.clients[c]; ok {
		c.connection.Close()
		close(c.egress)
		delete(m.clients, c)
	}
}
