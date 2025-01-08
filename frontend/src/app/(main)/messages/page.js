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
    <div className="flex h-[80vh]">
      {/* Main section for messages */}
      <main className="flex-1 p-4">
        <header className="border-b pb-2 mb-4">
          <h2 className="text-xl font-bold">Messages</h2>
        </header>

        {/* Messages container */}
        <div className="flex flex-col space-y-4 overflow-y-auto h-[calc(80vh-120px)]">
          <div className="bg-blue-100 p-3 rounded self-start max-w-md">
            Hello! How are you?
          </div>
          <div className="bg-green-100 p-3 rounded self-end max-w-md">
            I'm good, thanks! How about you?
          </div>
          <div className="bg-blue-100 p-3 rounded self-start max-w-md">
            Doing great! Let’s catch up soon.
          </div>
        </div>

        {/* Message input */}
        <div className="mt-4 flex items-center space-x-2">
          <input
            type="text"
            className="flex-1 border border-gray-300 rounded p-2"
            placeholder="Type your message..."
          />
          <button className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
            Send
          </button>
        </div>
      </main>

      {/* Sidebar for friends */}
      <aside className="w-1/4 bg-gray-100 p-4 border-l border-gray-200">
        <h2 className="text-xl font-bold mb-4">Friends</h2>
        <ul className="space-y-2">
          <li className="p-2 bg-gray-200 rounded hover:bg-gray-300">John Doe</li>
          <li className="p-2 bg-gray-200 rounded hover:bg-gray-300">Jane Smith</li>
          <li className="p-2 bg-gray-200 rounded hover:bg-gray-300">Alice Brown</li>
          {/* Add more friends here */}
        </ul>
      </aside>
    </div>
  );
}
