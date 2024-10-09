"use client";
import { useEffect, useState } from "react";
import FetchFromBackend from "@/lib/fetch";

export default function Group({ params }) {
  const { groupname } = params;
  const [groupData, setGroupData] = useState(null);
  const [posts, setPosts] = useState([]);
  const [isOwner, setIsOwner] = useState(false);
  const [loading, setLoading] = useState(true);

  // fetch group information
  useEffect(() => {
    async function loadData() {
      setLoading(true);
      console.log(`starting process for ${groupname}`);

      // fetch group information from groups table
      const response = await FetchFromBackend(`/group/${groupname}`, {
        method: "GET",
      });
      if (!response.ok) {
        throw new Error(`Failed to fetch ${groupname}`);
      }
      const data = await response.json();
      console.log(data.group);
      setGroupData(data.group);

      setLoading(false);
    }
    loadData();
  }, [groupname]);

  // TODO: check the group is created by own (compare login info and creatorId)
  // if the group's own, show invite user option

  // TODO: check the user is member of group
  // TODO: show join request (not a member)
  // TODO: show posts & events (member)

  // show loading message while the process is going on
  if (loading) {
    return <div>Loading...</div>;
  }

  // show message when the group is not found
  if (!groupData) {
    return <div>The group doesn't exist.</div>;
  }

  return (
    <div className="group-page">
      <div className="flex flex-row items-center">
        <div className="ml-2">
          <h1 className="text-2xl">{groupData.title}</h1>
          <p className="text-gray-600">{groupData.description}</p>
        </div>
      </div>
    </div>
  );
}
