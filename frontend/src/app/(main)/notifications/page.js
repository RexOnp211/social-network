"use client";

import FetchFromBackend from "@/lib/fetch";
import fetchCredential from "@/lib/fetchCredential";
import WsClient, { FollowRequest } from "@/lib/wsClient";
import { useState, useEffect, useRef } from "react";

export default function Notifications() {
  const [notifications, setNotifications] = useState([]);
  const ws = useRef(null);

  useEffect(() => {
    const load = async () => {
      try {
        const wsClient = await WsClient();
        ws.current = wsClient;

        ws.current.onmessage = (event) => {
          if (event.type === "follow_request") {
            setNotifications((prevNotifications) => [
              ...prevNotifications,
              event.data,
            ]);
          } else if (event.type === "follow_request_list") {
            setNotifications((prevNotifications) => [
              ...prevNotifications,
              ...JSON.parse(event.data),
            ]);
          }
        };
      } catch (error) {
        console.error(error);
      }
    };
    load();
  }, []);

  const handleAccept = (index) => {
    // Handle accept logic here
    setNotifications((prevNotifications) =>
      prevNotifications.filter((_, i) => i !== index),
    );
  };

  const handleDecline = (index) => {
    // Handle decline logic here
    setNotifications((prevNotifications) =>
      prevNotifications.filter((_, i) => i !== index),
    );
  };

  return (
    <div>
      <h1>Notifications</h1>
      <h1>{notifications}</h1>
      <ul>
        {notifications.map((notification, index) => (
          <li key={index}>
            {notification}
            <button onClick={() => handleAccept(index)}>Accept</button>
            <button onClick={() => handleDecline(index)}>Decline</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
