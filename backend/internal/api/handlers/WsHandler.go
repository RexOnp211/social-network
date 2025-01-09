package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var manager = NewManager()

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}

	token := r.Cookies()[0].Value
	userId, err := db.GetUserIDFromSession(token)
	if err != nil {
		log.Println("Error getting user id from session in ws", err)
		return
	}
	nickname := db.GetNicknameFromId(strconv.Itoa(userId))
	client := NewClient(conn, manager, userId, nickname)
	manager.AddClient(client)

	fmt.Println("these are the current clients", manager)

	go client.Read()
	go client.Write()
}

func SendFollowRequestHandler(event Event, c *Client) error {
	var followRequest helpers.FollowRequest
	if err := json.Unmarshal(event.Payload, &followRequest); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	fmt.Println("followrequest", followRequest)
	err := db.AddFollowRequestToDb(followRequest)
	if err != nil {
		return fmt.Errorf("users already follow eachother or sth went wrong: %v", err)
	}

	followRequestId, err := strconv.Atoi(followRequest.ToUserId)
	if err != nil {
		return fmt.Errorf("invalid user id: %v", err)
	}

	followRequests, err := db.GetFollowRequestsFromDb(followRequestId)
	if err != nil {
		return fmt.Errorf("error getting follow requests from db: %v", err)
	}

	followRequestsPayload, err := json.Marshal(followRequests)
	if err != nil {
		return fmt.Errorf("error marshalling follow requests: %v", err)
	}

	c.manager.RLock()
	defer c.manager.RUnlock()

	for client := range c.manager.clients {
		log.Println("client user ids", client.user.Id)
		if strconv.Itoa(client.user.Id) == followRequest.ToUserId {
			client.egress <- Event{
				Type:    SendFollowRequest,
				Payload: followRequestsPayload,
			}
		}
	}
	return nil
}

/*
Add event handlers down here
the functions should send event structs to the client's egress channel
*/

func GetFollowRequestsHandler(event Event, c *Client) error {
	fmt.Println("getting follow requests")
	followRequests, err := db.GetFollowRequestsFromDb(c.user.Id)
	if err != nil {
		return fmt.Errorf("error getting follow requests from db: %v", err)
	}

	followRequestsPayload, err := json.Marshal(followRequests)
	if err != nil {
		log.Println("Error marshaling follow request event:", err)
		return fmt.Errorf("error marshalling follow requests: %v", err)
	}

	c.egress <- Event{
		Type:    SendFollowRequest,
		Payload: followRequestsPayload,
	}

	return nil
}

func AcceptOrDeclineFollowRequest(event Event, c *Client) error {
	var followRequest helpers.FollowRequest
	if err := json.Unmarshal(event.Payload, &followRequest); err != nil {
		return fmt.Errorf("bad payload in request in AcceptFollowRequest: %v", err)
	}

	err2 := db.UpdateFollowRequestStatusDB(followRequest.FromUserId, followRequest.ToUserId, followRequest.FollowsBack)
	if err2 != nil {
		return err2
	}

	return nil
}

func GetChatMessagesWs(event Event, c *Client) error {

}
