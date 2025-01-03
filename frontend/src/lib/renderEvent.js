"use client";
import { useEffect, useState } from "react";
import formatDate from "@/lib/formatDate";
import FetchFromBackend from "@/lib/fetch";

export default function RenderEvent({ loggedInUsername, event }) {
  const [isGoing, setIsGoing] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      // fetch user's going status for the event
      const response = await FetchFromBackend(
        `/fetch_user_event_status/?username=${encodeURIComponent(
          loggedInUsername
        )}&event=${encodeURIComponent(event.Id)}`,
        {
          method: "GET",
        }
      );
      if (!response.ok) {
        console.log(`Failed to fetch ${loggedInUsername}`);
        return;
      }
      const status = await response.json();
      console.log("Data:", status);
      setIsGoing(status.going);
    };
    fetchData();
  }, []);

  // trigger when user change the event going status
  async function handleToggle() {
    console.log("toggle");
    try {
      const response = await FetchFromBackend("/update_event_status", {
        method: "POST",
        body: JSON.stringify({
          Nickname: loggedInUsername,
          EventId: event.Id,
          Going: !isGoing,
        }),
      });
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "An error occurred");
      }
      setIsGoing(!isGoing);
    } catch (error) {
      console.error(`Error updating status:`, error);
    }
  }

  return (
    <div>
      <div key={event.Id} className="bg-secondary p-4 rounded-lg m-4">
        <a className="flex flex-row items-center">
          {event.nickname || "loading..."}
        </a>
        <h1 className="text-xl font-bold">{event.title}</h1>
        <p>{event.description}</p>
        <p>Date of the event: {formatDate(event.eventDate)}</p>
        <span className="mr-2">Going to this event: </span>
        <div
          onClick={handleToggle}
          className={`relative inline-flex items-center cursor-pointer w-14 h-8 rounded-full transition-colors ${
            isGoing ? "bg-blue-500" : "bg-gray-300"
          }`}
        >
          <span
            className={`inline-block w-6 h-6 bg-white rounded-full transform transition-transform duration-300 ${
              isGoing ? "translate-x-6" : "translate-x-1"
            }`}
          />
        </div>
      </div>
    </div>
  );
}
