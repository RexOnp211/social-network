"use client";
import FetchFromBackend from "@/lib/fetch";
import { useEffect, useState } from "react";
import { useRef } from "react";
import WsClient, { ChatMessage } from "@/lib/wsClient";
import fetchCredential from "@/lib/fetchCredential";
import { FetchGroupPost } from "@/lib/fetchGroupPosts";
import Fetchnickname from "@/lib/fetchNickName";

export default function Messages() {
  const [userData, setUserData] = useState(null);
  const [friends, setFriends] = useState([]);
  const [selectedUser, setSelectedUser] = useState(1)
  const [selectedPrivateUser, setSelectedPrivateUser] = useState(1)
  const [messages, setMessages] = useState([])
  const [groupChats, setGroupChats] = useState([])
  const [nickname, setNickname] = useState({})
  const [chatType, setChatType] = useState("")
  const ws = useRef(null);
  const [loggedInUserId, setLoggedInUserId] = useState(null);

  useEffect(() => {
    const GetFriends = async () => {

    const user = await fetchCredential()

    const FolowersRes = await FetchFromBackend(`/followers/${user.username}`, {
      method: "GET",
      credentials: "include"
    })
    const followers = await FolowersRes.json()

    const followingRes = await FetchFromBackend(`/following/${user.username}`, {
      method: "GET",
      credentials: "include"
    })
    const following = await followingRes.json()
    const mergedUsers = [
      ...followers,
      ...following.filter(user => !followers.some(follower => follower.username === user.username))
      ];

      Promise.all(
        mergedUsers.map(user2 => {
          console.log("USER1", user.id, "USER2", user2.id);
          return FetchFromBackend("/chatId", {
            method: "POST",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              user1: user.id,
              user2: user2.id,
            }),
          })
            .then(res => res.text())
            .then(id => {
              user2.chatId = Number(id);
              return user2; // Return updated user2
            })
            .catch(err => {
              console.log(err);
              return user2; // Return user2 even if the request fails
            });
        })
      ).then(updatedUsers => {
        setFriends(updatedUsers)
      });
    }
   GetFriends()
  }, [])

useEffect(() => {
  const findNickname = async () => {
    const ids = [...new Set(messages.map((m) => m.fromUserId))]
    const nicknameMap = {};

    for (const userId of ids) {
      try {
        const res = await Fetchnickname(userId)
        const textData = await res.text()
        nicknameMap[userId] = textData
      } catch(error) {
        console.error("Error fetching nicknames in messages", error)
      }
    }
    setNickname(nicknameMap)
  }
  if (messages.length > 0) {
    findNickname()
  }
}, [messages])

  useEffect(() => {
    const Groups = async () => {
      const user = await fetchCredential()
      const membersRes = await FetchFromBackend(`/fetch_memberships/${user.username}`, {
        method: "GET",
        credentials: "include"
      })
      const members = await membersRes.json()
      const { invitedMemberships, approvedMemberships} =
      members.memberships.reduce(
        (acc, membership) => {
          if (membership.status === "invited") {
            acc.invitedMemberships.push(membership)
          } else if (membership.status === "approved") {
            acc.approvedMemberships.push(membership)
          }
          return acc
        },
        {invitedMemberships: [], approvedMemberships: []}
      )

      const GroupsRes = await FetchFromBackend(`/groups`, {
        method: "GET",
        credentials: "include"
      })
      const Groups = await GroupsRes.json()

      const userMadeGroups = Groups.groups.filter((group) => {
        return group.creatorName === user.username
      })
      console.log("USER MAD GROUPS", userMadeGroups)
      approvedMemberships.push(...userMadeGroups)
      console.log("MEMBERS", approvedMemberships)
      setGroupChats(approvedMemberships)
    }
    Groups()
  }, [])

  const switchChat = async (e) => {
    setSelectedUser(Number(e.target.id))
    setChatType("group")
    console.log("TARGET", e.target.id)
    console.log("SELECTED USER", selectedUser)
    // add logic here when backend part is done
    fetchMessages("group", Number(e.target.id))
  }

  const switchPrivateChat = async (e) => {
    setSelectedPrivateUser(Number(e.target.id))
    setChatType("user")
    console.log("TARGET", e.target.id)
    console.log("SELECTED USER", selectedPrivateUser)
    // add logic here when backend part is done
    fetchMessages("user", Number(e.target.id))
  }

  const sendMessage = async (e) => {
    e.preventDefault();
    if (ws.current) {
      const user = await fetchCredential();
      const otherUserId = selectedUser
      const form = new FormData(e.target)
      const content = form.get("message")
      console.log("form message", content)
      ws.current.send(
        JSON.stringify({
          type: "message_send",
          payload: new ChatMessage(
             chatType, user.id, otherUserId, content
          )
        })
      );
    }
    
  };

  useEffect(() => {
    const load = async () => {
      const wsClient = await WsClient();
      ws.current = wsClient;

      ws.current.onmessage = (event) => {
        const eventData = event.data
        console.log("EVENT", event)
        const parsedData = JSON.parse(eventData)
        console.log("eventData", parsedData)
        console.log("type", parsedData.type)
        if (parsedData.type === "messages") {
          console.log("true")
          setMessages(parsedData.payload)
        }
      };
    };
    load();
  }, []);

  const fetchMessages = async (type, chatId) => {
    console.log('TYPE', type, "CHATID", chatId)
    try {
      ws.current.send(
        JSON.stringify({
          type: "get_chat_messages",
          payload: {chatType: type, groupId: chatId}
        })
      )
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    const fetchLoggedInUser = async () => {
      const usr = await fetchCredential(); // Assumes this fetches the logged-in user's data
      setLoggedInUserId(usr.username);
    };
    fetchLoggedInUser();
  }, []);

  return (
    <div className="flex h-[80vh]">
      {/* Main section for messages */}
      <main className="flex-1 p-4">
        <header className="border-b pb-2 mb-4">
          <h2 className="text-xl font-bold">Messages</h2>
        </header>

        {/* Messages container */}
        <div className="flex flex-col space-y-4 overflow-y-auto h-[calc(80vh-120px)]">
        {messages.length === 0 ? (
  <div className="bg-blue-100 p-3 rounded self-start max-w-md">
    click on a friend on the right to start messaging
  </div>
) : (
  messages.map((message, index) => (
        <div
          key={index}
          className={`p-3 rounded max-w-md ${
            loggedInUserId === nickname[message.fromUserId]
              ? "bg-green-100 self-end"
              : "bg-blue-100 self-start"
          }`}
        >
          <p>{nickname[message.fromUserId]}</p>
          <p>{message.content}</p>
        </div>
  ))
)}

          
          {/*<div className="bg-blue-100 p-3 rounded self-start max-w-md">Hello! How are you?</div>
           <div className="bg-green-100 p-3 rounded self-end max-w-md">
            I'm good, thanks! How about you?
          </div>
          <div className="bg-blue-100 p-3 rounded self-start max-w-md">
            Doing great! Let’s catch up soon.
          </div> */}
        </div>

        {/* Message input */}
        <form id="msg" onSubmit={sendMessage}>
        <div className="mt-4 flex items-center space-x-2">
          <input
            name="message"
            id="message"
            type="text"
            className="flex-1 border border-gray-300 rounded p-2"
            placeholder="Type your message..."
          />
          <div></div>
          <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
            Send
          </button>
        </div>
        </form>
      </main>

      {/* Sidebar for friends */}
      <aside className="w-1/4 bg-gray-100 p-4 border-l border-gray-200">
        <h2 className="text-xl font-bold mb-4">Friends</h2>
        <ul className="space-y-2">
          {friends.length === 0 ? (
            <li className="p-2 bg-gray-200 rounded hover:bg-gray-300">No Users Found</li>
          ) : (
            friends.map((user) => (
              <li className="p-2 bg-gray-200 rounded hover:bg-gray-300" key={user.id}>
              <button key={user.id} id={user.chatId} onClick={switchPrivateChat} >
                {user.nickname}
              </button>
              </li>
            ))
          )}
        </ul>
        <h2 className="text-xl font-bold mb-4">Groups</h2>
        <ul className="space-y-2">
          {groupChats.length === 0 ? (
            <li className="p-2 bg-gray-200 rounded hover:bg-gray-300">No Groups found</li>
          ) : (
            groupChats.map((group, index) => (
              <li className="p-2 bg-gray-200 rounded hover:bg-gray-300" key={index}>
              <button key={index} id={group.chatId} onClick={switchChat} >
                {group.title}
              </button>
              </li>
            ))
          )
        
        }
        </ul>
      </aside>
    </div>
  );
}
/*

return (
  <div class="flex flex-col justify-center border border-gray-300 p-6 rounded-lg shadow-md bg-white">
    <div>
      <h1>Chat</h1>
      {messages.length === 0 ? (
        <p>No messages found, start one!</p>
      ) : (
        messages.map((message) => (
          <a>message</a>
        ))
      )}
      <form>
        <input type="text" id="message"></input>
        <button class="bg-accent w-full text-white rounded-lg p-3 transition-colors hover:bg-accentDark" onClick={sendMessage}>Send</button>
      </form>
    <div className="w-64 h-auto m-3 bg-primary text-txtColor flex flex-col p-2 rounded-lg shadow-lg absolute inset-y-0 right-0">
    <nav className="x-auto flex flex-col space-y-4 overflow-y-auto">
      </div>
    </nav>
    </div>
  </div>
  </div>
);

*/