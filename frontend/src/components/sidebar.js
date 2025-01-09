"use client";

import FetchFromBackend from "@/lib/fetch";
import fetchCredential from "@/lib/fetchCredential";
import { RSC_ACTION_CLIENT_WRAPPER_ALIAS } from "next/dist/lib/constants";
import { useEffect, useState } from "react";

export default function SideBar() {
  const [showMoreFollowing, setShowMoreFollowing] = useState(false);
  const [following, setFollowing] = useState([])
  
  useEffect(() => {
    const getFollowing = async () => {
      const user = await fetchCredential()
      const res = await FetchFromBackend(`/following/${user.username}`, {
        method: "GET",
        credentials: "include",
      })
      const users = await res.json()
      setFollowing(users)
    }
    getFollowing()
  }, [])

  return (
    <div className="w-64 h-[87vh] m-3 bg-primary text-txtColor flex flex-col p-2 rounded-lg shadow-lg">
      <nav className="flex flex-col space-y-4 overflow-y-auto">
        <div className="transition-colors hover:bg-secondary ease-in p-2 rounded">
          <h1 className="text-lg text-accent font-bold">Following</h1>
          <RenderList
            items={following}
            showMoreState={showMoreFollowing}
            setShowMoreState={setShowMoreFollowing}
            type="following"
          />
        </div>
        <div className="transition-colors hover:bg-secondary ease-in p-2 rounded">
          <a href="/group" className="text-txtColor hover:underline">
            <h1 className="text-lg text-accent font-bold">Create Group</h1>
          </a>
        </div>
      </nav>
    </div>
  );
}

const RenderList = ({ items, showMoreState, setShowMoreState, type }) => {
  console.log("items", items)
  const validItems = Array.isArray(items) ? items : []
  let showMoreitems = showMoreState ? validItems: validItems.slice(0,4)
  return (
    <>
      <ul className="list-disc pl-5 marker:text-txtColor">
        {showMoreitems.flatMap((item, index) => (
          <li key={index}>
            <a
              href={`/${type === "event" ? "events" : type === "group" ? "groups" : "profile"}/${item.nickname}`}
              className="text-txtColor hover:underline"
            >
              {(type !== "event" && type !== "group") ? item.nickname : item}
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
