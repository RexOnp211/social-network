"use client";
import { useState, useEffect } from "react";
import FetchFromBackend from "@/lib/fetch";

export default function GroupEvent({
  loggedInUserID,
  loggedInUsername,
  groupTitle,
  showPopup,
  onEventSubmit,
}) {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [eventDate, setEventDate] = useState("");

  async function handleSubmit(e) {
    e.preventDefault();

    try {
      const form = new FormData(e.target);
      form.append("groupname", groupTitle);
      form.append("user_id", loggedInUserID);
      form.append("nickname", loggedInUsername);

      console.log([...form.entries()]);

      const response = await FetchFromBackend("/create_group_event", {
        method: "POST",
        body: form,
      });
      if (!response.ok) {
        const errorText = await response.text();
        console.log(errorText);
        showPopup(true, "Event date cannot be in the past", 5000);
        return;
      }

      // clear form
      setTitle("");
      setDescription("");
      setEventDate("");

      showPopup(false, "Event created successfully!", 3000);

      // update posts
      if (onEventSubmit) {
        setTimeout(async () => {
          await onEventSubmit(groupTitle);
        }, 1000);
      }
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <div className="mt-4 mb-4">
      <div
        className={
          "transition-max-height duration-500 ease-in-out overflow-hidden max-h-screen w-full"
        }
      >
        <form
          id="CreateGroupEvent"
          className="flex flex-col justify-center border border-gray-300 p-6 rounded-lg shadow-md bg-white"
          onSubmit={handleSubmit}
          encType="multipart/form-data"
        >
          {/* Event Title */}
          <label
            htmlFor="eventTitle"
            className="mb-2 text-lg font-semibold text-gray-700"
          >
            Event Title <span className="text-red-500">*</span>
          </label>
          <input
            type="text"
            id="eventTitle"
            name="eventTitle"
            placeholder="Enter event title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
            className="mb-4 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-accent focus:bg-primary"
          />

          {/* Event Description */}
          <label
            htmlFor="eventDescription"
            className="mb-2 text-lg font-semibold text-gray-700"
          >
            Event Body <span className="text-red-500">*</span>
          </label>
          <textarea
            id="eventDescription"
            name="eventDescription"
            placeholder="Enter event description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            required
            className="mb-4 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-accent focus:bg-primary"
          ></textarea>

          {/* Event Date */}
          <label
            htmlFor="eventDate"
            className="mb-2 text-lg font-semibold text-gray-700"
          >
            Event Date <span className="text-red-500">*</span>
          </label>
          <input
            type="datetime-local"
            id="eventDate"
            name="eventDate"
            value={eventDate}
            onChange={(e) => setEventDate(e.target.value)}
            required
            className="mb-4 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-accent focus:bg-primary"
          />

          <button
            type="submit"
            className="bg-accent w-full text-white rounded-lg p-3 transition-colors hover:bg-accentDark"
          >
            Create Event
          </button>
        </form>
      </div>
    </div>
  );
}
