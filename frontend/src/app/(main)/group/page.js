"use client";
import { useEffect, useState } from "react";
import FetchFromBackend from "@/lib/fetch";
import RenderList from "@/components/renderList";
import GroupInvitation from "@/components/groupInvitation";

export default function GroupMenu({ params }) {
  const [username, setUsername] = useState(null);
  const [loading, setLoading] = useState(true);
  const [groups, setGroups] = useState([]);
  const [groupsUserMade, setGroupUserMade] = useState([]);
  const [groupsUserJoined, setGroupUserJoined] = useState([]);
  const [invitationExist, setInvitationExist] = useState(false);
  const [showMoreGroups, setShowMoreGroups] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [popupMessage, setPopupMessage] = useState("");
  const [showPopup, setShowPopup] = useState(false);
  const [storedUsername, setStoredUsername] = useState(
    localStorage.getItem("user")
  );
  const [approvedInvs, setApprovedInvs] = useState([]);

  useEffect(() => {
    // fetch login user name
    console.log("Loaded username from localStorage:", storedUsername);
    setUsername(storedUsername);
    console.log(storedUsername);

    fetchData(storedUsername);
  }, []);

  const fetchData = async (storedUsername) => {
    // fetch group invitations for the user
    await loadAllInvitations(storedUsername);
    // fetch all groups
    await loadAllGroups();
  };

  // fetch group invitations for the user
  async function loadAllInvitations(storedUsername) {
    console.log("loadAllInvitations");
    try {
      const response = await FetchFromBackend(
        `/group_member/${storedUsername}`,
        {
          method: "GET",
        }
      );
      if (!response.ok) {
        console.error(
          "Failed to fetch group_member, response status:",
          response.status
        );
        throw new Error(`Failed to fetch group_member`);
      }

      const data = await response.json();
      console.log("Data received in loadAllInvitations:", data);

      // Normalize the field names to ensure consistency
      const normalizedData = data.groupmembers.map((invitation) => ({
        ...invitation,
        status: invitation.Status || invitation.status,
      }));

      // Filtering invitations with status "invited"
      const { invitedInvitations, approvedInvitations } = normalizedData.reduce(
        (acc, invitation) => {
          if (invitation.status === "invited") {
            acc.invitedInvitations.push(invitation);
          } else if (invitation.status === "approved") {
            acc.approvedInvitations.push(invitation);
          }
          return acc;
        },
        { invitedInvitations: [], approvedInvitations: [] }
      );

      console.log("Invited Invitations:", invitedInvitations);
      console.log("Approved Invitations:", approvedInvitations);
      setInvitationExist(invitedInvitations);

      // filter groups user joined
      setApprovedInvs(approvedInvitations);
    } catch (error) {
      console.error("Error in loadAllInvitations:", error);
    }
  }

  // fetch all groups
  async function loadAllGroups() {
    setLoading(true);

    const response = await FetchFromBackend(`/groups`, {
      method: "GET",
    });
    if (!response.ok) {
      throw new Error(`Failed to fetch groups`);
    }
    const data = await response.json();
    console.log(data.groups);
    setGroups(data.groups);

    // Filter  groups the user made
    const userMadeGroups = data.groups.filter((group) => {
      return group.creatorName === storedUsername;
    });
    setGroupUserMade(userMadeGroups);

    const approvedTitles = approvedInvs.map((invitation) => invitation.title);
    const filteredGroups = data.groups.filter((group) =>
      approvedTitles.includes(group.title)
    );
    setGroupUserJoined(filteredGroups);

    setLoading(false);
  }

  // create new group
  const OnSubmit = async (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    console.log(username);
    formData.append("user_id", username);

    try {
      const response = await FetchFromBackend("/creategroup", {
        method: "POST",
        body: formData,
      });
      if (!response.ok) {
        throw new Error("Failed to create group");
      }

      // clear form
      setTitle("");
      setDescription("");

      // show successful popup message
      const groupname = formData.get("title");
      const message = `Group ${groupname} successfully created.`;
      setPopupMessage(message);
      setShowPopup(true);

      // update display
      await fetchData(storedUsername);

      // Hide the message after 3 seconds
      setTimeout(() => {
        setShowPopup(false);
      }, 3000);
    } catch (error) {
      console.error("Error submitting form:", error);
    }
  };

  return (
    <div className="groupMenu-page w-full">
      {invitationExist && invitationExist.length != 0 && (
        <GroupInvitation
          invitationExist={invitationExist}
          username={username}
          onAcceptOrReject={() => fetchData(storedUsername)}
        />
      )}
      <div className="flex flex-col items-start w-full">
        <div className="w-full">
          <h2 className="text-lg text-accent font-bold">Create New Group</h2>
          <form onSubmit={OnSubmit} className="w-full">
            <div className="w-full">
              <label htmlFor="title" className="block mt-2">
                Title <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                id="title"
                name="title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                required
                className="p-2 rounded-lg w-full max-w-full"
              />
            </div>
            <div className="w-full">
              <label htmlFor="description" className="block mt-2">
                Description <span className="text-red-500">*</span>
              </label>
              <textarea
                id="description"
                name="description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                required
                className="p-2 rounded-lg w-full"
              />
            </div>
            <button
              type="submit"
              className="mt-2 transition-colors ease-in hover:bg-accentDark bg-accent text-white rounded-lg p-2 w-full"
            >
              Create
            </button>
          </form>
          <h2 className="mt-8 text-lg text-accent font-bold">
            Groups You Made
          </h2>
          {loading ? (
            <p>Loading...</p>
          ) : groupsUserMade.length === 0 ? (
            <p>No groups yet</p>
          ) : (
            <RenderList
              items={groupsUserMade.map((group) => group.title)}
              showMoreState={showMoreGroups}
              setShowMoreState={setShowMoreGroups}
              type="group"
            />
          )}
        </div>
      </div>
      <div className="w-full">
        <h2 className="mt-8 text-lg text-accent font-bold">
          Groups You Joined
        </h2>
        {loading ? (
          <p>Loading...</p>
        ) : approvedInvs.length === 0 ? (
          <p>No groups yet</p>
        ) : (
          <RenderList
            items={approvedInvs.map((group) => group.title)}
            showMoreState={showMoreGroups}
            setShowMoreState={setShowMoreGroups}
            type="group"
          />
        )}
      </div>
      <div className="w-full">
        <h2 className="mt-8 text-lg text-accent font-bold">
          Browse All Groups
        </h2>
        {loading ? (
          <p>Loading...</p>
        ) : groups.length === 0 ? (
          <p>No groups yet</p>
        ) : (
          <RenderList
            items={groups.map((group) => group.title)}
            showMoreState={showMoreGroups}
            setShowMoreState={setShowMoreGroups}
            type="group"
          />
        )}
      </div>

      {showPopup && (
        <div className="absolute bottom-0 left-0 p-4 mb-10 ml-10 bg-orange-100 text-black rounded shadow-lg">
          {popupMessage}
        </div>
      )}
    </div>
  );
}
