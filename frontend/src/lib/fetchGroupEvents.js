"use client";
import FetchFromBackend from "@/lib/fetch";

export default async function FetchGroupEvents(groupname) {
  console.log("fetching events");
  const eventResponse = await FetchFromBackend(
    `/fetch_group_events/${groupname}`,
    {
      method: "GET",
    }
  );
  if (!eventResponse.ok) {
    throw new Error(`Failed to fetch group events ${groupname}`);
  }
  const events = await eventResponse.json();
  console.log("events", events);
  return events.data;
}
