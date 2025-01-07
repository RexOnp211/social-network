"use client";
import FetchFromBackend from "@/lib/fetch";
import { useEffect, useState } from "react";
import { useRef } from "react";
import WsClient from "@/lib/wsClient";
import fetchCredential from "@/lib/fetchCredential";

export default function Messages() {
  const [userData, setUserData] = useState(null);

  const users = [];
  const messages = [];
  const ws = useRef(null);

  const sendMessage = async (e) => {
    if (ws.current) {
      const user = await fetchCredential();
      const otherUserId = await userData.id;

      ws.current.send(
        JSON.stringify({
          type: "message_send",
          payload: new MessageSend(
            user.id, otherUserId, content
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
        if (event.type === "message_send") {
          alert(`Message received: ${event.data}`);
        }
      };
    };
    load();
  }, []);

  const fetchMessages = async (e) => {
    try {
      const path = ``;
      const res = await FetchFromBackend(path, 
        {method: "GET"}
      )
    } catch (error) {
      console.error(error);
    }
  }

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
        <div className="transition-colors hover:bg-secondary ease-in p-2 rounded">
          <h1 className="text-lg text-accent font-bold">Users</h1>
          {users.length === 0 ? (
          <p>No users found</p>
        ) : (
          users.map((user) => (
            <a>user</a>
          ))
        )}
        </div>
      </nav>
      </div>
    </div>
    </div>
  );
}
