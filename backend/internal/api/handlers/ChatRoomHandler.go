package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"
)

/* use this function to send any messages to the chat room (group_id, from_which_user, message_content) */
func PrivateMessageHandler(event Event, c *Client) error {
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

/* use this function to add a new chatroom entry */
func ChatRoomHandler(event Event, c *Client) error {
	var chatRoom helpers.ChatRoom
	if err := json.Unmarshal(event.Payload, &chatRoom); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	fmt.Println("ChatRoomHandler print ", chatRoom)
	err := db.AddChatRoomIntoDb(chatRoom)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

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
