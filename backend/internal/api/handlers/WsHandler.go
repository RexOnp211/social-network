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

func AddMessageIntoDb(event Event, c *Client) error {
	ChatMsg := struct {
		ChatType   string `json:"chatType"`
		GroupId    int    `json:"groupId"`
		FromUserId int    `json:"fromUserId"`
		Content    string `json:"content"`
	}{}

	fmt.Println("This is the event Payload that i just sent to backend", event.Payload)

	err := json.Unmarshal(event.Payload, &ChatMsg)
	if err != nil {
		fmt.Println("ERROR IN UNMARSHALING?", err)
		return fmt.Errorf("bad payload in WS GetChatMessagesWS: %v", err)
	}
	fmt.Println("this is the payload after unmarshaling it in getgroupchatmessagesws", ChatMsg)
	if ChatMsg.ChatType == "group" {
		err2 := db.AddChatMessageIntoDb(ChatMsg.GroupId, ChatMsg.FromUserId, ChatMsg.Content)
		if err2 != nil {
			fmt.Println("Error adding message to db", err2)
			return fmt.Errorf("error in GetGroupCatMessagesWs: %v", err)
		}
	} else if ChatMsg.ChatType == "user" {
		err2 := db.AddPrivateChatMessageIntoDb(ChatMsg.GroupId, ChatMsg.FromUserId, ChatMsg.Content)
		if err2 != nil {
			fmt.Println("Error adding private message to db")
		}
	}

	GetChatMessagesWs(event, c)

	return nil
}

func GetChatMessagesWs(event Event, c *Client) error {
	msg := struct {
		MsgType string `json:"chatType"`
		GroupId int    `json:"groupId"`
	}{}

	err := json.Unmarshal(event.Payload, &msg)
	if err != nil {
		fmt.Println("Error unmarshaling in GetChatMessages", err)
		return fmt.Errorf("bad payload in GetChatMessagesWs: %v", err)
	}

	fmt.Println("MSG", msg)
	var messagePayload []byte
	membersIds := []int{}

	switch msg.MsgType {
	case "user":
		// Fetch private messages
		privateMessages, err := db.GetChatMessagesFromDb(msg.GroupId) // Returns []PrivateMessage
		if err != nil {
			fmt.Println("Error getting private messages from DB", err)
			return fmt.Errorf("Error getting private messages: %v", err)
		}
		messagePayload, err = json.Marshal(privateMessages)
		if err != nil {
			fmt.Println("Error marshaling private messages in GetChatMessages", err)
			return fmt.Errorf("Error marshaling private messages: %v", err)
		}
		// Add private chat members
		fromUserIds := map[int]struct{}{}
		for _, pm := range privateMessages {
			fromUserIds[pm.FromUserId] = struct{}{}
		}
		for userId := range fromUserIds {
			membersIds = append(membersIds, userId)
		}

	case "group":
		// Fetch group chat messages
		chatMessages, err := db.LoadChatRoomMessages(msg.GroupId) // Returns []ChatMessage
		if err != nil {
			fmt.Println("Error getting group chat messages from DB", err)
			return fmt.Errorf("Error getting group chat messages: %v", err)
		}
		messagePayload, err = json.Marshal(chatMessages)
		if err != nil {
			fmt.Println("Error marshaling group chat messages in GetChatMessages", err)
			return fmt.Errorf("Error marshaling group chat messages: %v", err)
		}

		// Fetch members in the group
		title, creatorName, err := db.GetGroupWithChatId(msg.GroupId)
		if err != nil {
			fmt.Println("Error getting group title and creator", err)
			return fmt.Errorf("Error getting group title and creator: %v", err)
		}
		creator, err := db.GetUserFromDb(creatorName)
		if err != nil {
			fmt.Println("Error getting creator info", err)
			return fmt.Errorf("Error getting creator info: %v", err)
		}
		membersIds = append(membersIds, creator.Id)

		groupMembers, err := db.GetApprovedMembershipsFromDb(title)
		if err != nil {
			fmt.Println("Error getting approved memberships from DB", err)
			return fmt.Errorf("Error getting approved memberships: %v", err)
		}
		for _, member := range groupMembers {
			user, err := db.GetUserFromDb(member.Username)
			if err != nil {
				fmt.Println("Error getting user info", err)
				return fmt.Errorf("Error getting user info: %v", err)
			}
			membersIds = append(membersIds, user.Id)
		}

	default:
		return fmt.Errorf("unknown MsgType: %s", msg.MsgType)
	}

	// Send the payload to the relevant clients
	for client := range c.manager.clients {
		for _, userId := range membersIds {
			if client.user.Id == userId {
				client.egress <- Event{
					Type:    "messages",
					Payload: messagePayload,
				}
			}
		}
	}

	return nil
}

func InviteMemberWs(event Event, c *Client) error {
	invitation := struct {
		Groupname string `json:"groupname"`
		Username  string `json:"username"`
	}{}
	err := json.Unmarshal(event.Payload, &invitation)
	if err != nil {
		log.Println("Failed to decode JSON")
		return err
	}
	log.Println("inviting member...", invitation)
	chatId, err := db.GetChatIdFromGroup(invitation.Groupname)
	if err != nil {
		fmt.Println("Error getting chatid from group", err)
		return err
	}
	errorMessage, isError := db.InviteMemberDB(invitation.Groupname, invitation.Username, chatId)
	fmt.Println("TEST", errorMessage, isError)
	if isError {
		if errorMessage != "" {
			fmt.Println("ERROR in INVIVTE MEMBER", errorMessage)
		}
	}
	for client := range c.manager.clients {
		if client.user.Nickname == invitation.Username {
			client.egress <- Event{
				Type:    "group_invite",
				Payload: event.Payload,
			}
		}
	}
	return nil
}
