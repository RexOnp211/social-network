"use client";

import { useState } from "react";

export default function SideBar() {
  const [showMoreFollowers, setShowMoreFollowers] = useState(false);
  const [showMoreFollowing, setShowMoreFollowing] = useState(false);
  const [showMoreGroups, setShowMoreGroups] = useState(false);
  const [showMoreEvents, setShowMoreEvents] = useState(false);

  const followers = [
    "Follower 1",
    "Follower 2",
    "Follower 3",
    "Follower 4",
    "Follower 5",
  ];
  const following = [
    "Following 1",
    "Following 2",
    "Following 3",
    "Following 4",
    "Following 5",
  ];
  const groups = ["Group 1", "Group 2", "Group 3", "Group 4", "Group 5"];
  const events = [
    "Event 1",
    "Event 2",
    "Event 3",
    "Event 4",
    "Event 5",
    "Event 6",
    "Event 7",
    "Event 8",
    "Event 9",
    "Event 10",
  ];
  return (
    <div className="w-64 h-[87vh] m-3 bg-primary text-txtColor flex flex-col p-2 rounded-lg shadow-lg">
      <nav className="flex flex-col space-y-4 overflow-y-auto">
        <div className="transition-colors hover:bg-secondary ease-in p-2 rounded">
          <h1 className="text-lg text-accent font-bold">Followers</h1>
          {renderList(
            followers,
            showMoreFollowers,
            setShowMoreFollowers,
            "follower",
          )}
        </div>
        <div className="transition-colors hover:bg-secondary ease-in p-2 rounded">
          <h1 className="text-lg text-accent font-bold">Following</h1>
          {renderList(
            following,
            showMoreFollowing,
            setShowMoreFollowing,
            "following",
          )}
        </div>
        <div className="transition-colors hover:bg-secondary ease-in p-2 rounded">
          <h1 className="text-lg text-accent font-bold">Browse Groups</h1>
          {renderList(groups, showMoreGroups, setShowMoreGroups, "group")}
        </div>
        <div className="transition-colors hover:bg-secondary ease-in p-2 rounded">
          <a href="/createGroup" className="text-txtColor hover:underline">
            <h1 className="text-lg text-accent font-bold">Create Group</h1>
          </a>
        </div>
        <div className="transition-colors hover:bg-secondary ease-in p-2 rounded">
          <h1 className="text-lg text-accent font-bold">Upcoming Events</h1>
          {renderList(events, showMoreEvents, setShowMoreEvents, "event")}
        </div>
      </nav>
    </div>
  );
}

const renderList = (items, showMoreState, setShowMoreState, type) => {
  return (
    <>
      <ul className="list-disc pl-5 marker:text-txtColor">
        {(showMoreState ? items : items.slice(0, 4)).map((item, index) => (
          <li key={index}>
            <a
              href={`/${type === "event" ? "events" : type === "group" ? "groups" : "profile"}/${type}${index + 1}`}
              className="text-txtColor hover:underline"
            >
              {item}
            </a>
          </li>
        ))}
      </ul>
      {items.length > 4 && showMore(showMoreState, setShowMoreState)}
    </>
  );
};

const showMore = (state, setState) => {
  return (
    <button
      onClick={() => setState(!state)}
      className="hover:underline flex items-center"
    >
      {state ? "Show Less" : "Show More"}
      <span className={`ml-1 transform ${state ? "rotate-180" : ""}`}>▼</span>
    </button>
  );
};
