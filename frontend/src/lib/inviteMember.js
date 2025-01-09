import FetchFromBackend from "@/lib/fetch";
import WsClient from "./wsClient";
import { useRef } from "react";

export default async function InviteMember(groupname, username) {
  console.log(groupname, username);
  try {
    const wsClient = await WsClient();
    const ws = await wsClient;

    // Ensure WebSocket is open before sending a message
    if (ws.readyState === WebSocket.CONNECTING) {
      await new Promise((resolve) => {
        ws.addEventListener("open", resolve, { once: true });
      });
    }

    if (ws.readyState === WebSocket.OPEN) {
      ws.send(
        JSON.stringify({
          type: "group_invite",
          payload: {
            groupname: groupname,
            username: username,
          },
        })
      );
    } else {
      throw new Error("WebSocket is not open");
    }
  } catch (error) {
    return `An error occurred: ${error.message}`;
  }

  return "";
}

