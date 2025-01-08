package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"
)

// func ChatRoomHandler(w http.ResponseWriter, r *http.Request) {
// 	var event Event
// 	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	client := &Client{}

// 	if err := CreateChatRoom(event, client); err != nil {
// 		http.Error(w, fmt.Sprintf("Handler error: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	_, _ = w.Write([]byte("Chat room created successfully"))
// }

func PrivateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	client := &Client{}

	if err := CreateMessage(event, client); err != nil {
		http.Error(w, fmt.Sprintf("Handler error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Private message created successfully"))
}

/* use this function to send any messages to the chat room (group_id, from_which_user, message_content) */
func CreateMessage(event Event, c *Client) error {
	var privateMessage helpers.PrivateMessage
	if err := json.Unmarshal(event.Payload, &privateMessage); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	fmt.Println("PrivateMessageHandler print ", privateMessage)
	err := db.AddChatMessageIntoDb(privateMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// /* use this function to add a new chatroom entry */
// func CreateChatRoom(event Event, c *Client) error {
// 	var chatRoom helpers.ChatRoom
// 	if err := json.Unmarshal(event.Payload, &chatRoom); err != nil {
// 		return fmt.Errorf("bad payload in request: %v", err)
// 	}

// 	fmt.Println("ChatRoomHandler print ", chatRoom)
// 	err := db.AddChatRoomIntoDb(chatRoom)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	return nil
// }

/* use this function to add users to a chat room (group_id, user) */
func AddUserToChatRoom(event Event, c *Client) error {
	var chatRoomMembers helpers.ChatRoomMembers
	if err := json.Unmarshal(event.Payload, &chatRoomMembers); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	fmt.Println("ChatRoomHandler print ", chatRoomMembers)
	err := db.AddUserIntoChatRoom(chatRoomMembers)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
