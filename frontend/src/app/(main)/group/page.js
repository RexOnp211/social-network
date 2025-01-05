"use client";
import { useEffect, useState } from "react";
import FetchFromBackend from "@/lib/fetch";
import RenderList from "@/components/renderList";
import GroupInvitation from "@/components/groupInvitation";
import GroupRequests from "@/components/groupRequests";
import Popup from "@/components/popup";

export default function GroupMenu() {
  const [loggedInUsername, setLoggedInUsername] = useState(
    localStorage.getItem("user")
  );
  const [invitations, setInvitations] = useState(false);
  const [joinedGroups, setJoinedGroups] = useState([]);
  const [requests, setRequests] = useState([]);
  const [groups, setGroups] = useState([]);
  const [groupsUserMade, setGroupUserMade] = useState([]);
  const [showMoreUserMadeGroups, setShowMoreUserMadeGroups] = useState(false);
  const [showMoreUserJoinedGroups, setShowMoreUserJoinedGroups] =
    useState(false);
  const [showMoreGroups, setShowMoreGroups] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [membershipsLoading, setMembershipsLoading] = useState(true);
  const [requestsLoading, setRequestsLoading] = useState(true);

  // load necessary data when the page is accessed
  useEffect(() => {
    (async () => {
      await fetchMemberships(loggedInUsername);
      await fetchGroups(loggedInUsername);
    })();
  }, []);

  // fetch group memberships for the user
  async function fetchMemberships(loggedInUsername) {
    try {
      const response = await FetchFromBackend(
        `/fetch_memberships/${loggedInUsername}`,
        {
          method: "GET",
        }
      );
      if (!response.ok) {
        throw new Error(`Failed to fetch memberships: ${response.status}`);
      }
      const data = await response.json();
      console.log("Data received in memberships:", data);

      // Filtering invitations with status "invited"
      const { invitedMemberships, approvedMemberships } =
        data.memberships.reduce(
          (acc, membership) => {
            if (membership.status === "invited") {
              acc.invitedMemberships.push(membership);
            } else if (membership.status === "approved") {
              acc.approvedMemberships.push(membership);
            }
            return acc;
          },
          { invitedMemberships: [], approvedMemberships: [] }
        );

      console.log("invitedMemberships:", invitedMemberships);
      console.log("approvedMemberships:", approvedMemberships);
      setInvitations(invitedMemberships);
      setMembershipsLoading(false);

      // filter groups user joined
      setJoinedGroups(approvedMemberships);
    } catch (error) {
      console.error("Error in fetchMemberships:", error);
    }
  }

  // fetch all groups
  async function fetchGroups(loggedInUsername) {
    const response = await FetchFromBackend(`/groups`, {
      method: "GET",
    });
    if (!response.ok) {
      throw new Error(`Failed to fetch groups`);
    }
    const data = await response.json();
    console.log(data.groups);
    setGroups(data.groups);

    // Filter groups the user made
    const userMadeGroups = data.groups.filter((group) => {
      return group.creatorName === loggedInUsername;
    });
    setGroupUserMade(userMadeGroups);
    console.log(userMadeGroups);

    FetchRequests(userMadeGroups);
  }

  // fetch group request for your group
  async function FetchRequests(userMadeGroups) {
    console.log("fetch_your_requests");
    try {
      const response = await FetchFromBackend(`/fetch_your_requests`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userMadeGroups),
      });
      if (!response.ok) {
        throw new Error(`Failed to fetch_your_requests`);
      }
      const data = await response.json();
      console.log("fetch_your_requests:", data);

      setRequests(data);
      setRequestsLoading(false);
    } catch (error) {
      console.error("Error in fetch_your_requests:", error);
    }
  }

  // create new group
  const OnSubmit = async (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    console.log(loggedInUsername);
    formData.append("user", loggedInUsername);

    try {
      const response = await FetchFromBackend("/create_group", {
        method: "POST",
        body: formData,
      });
      if (!response.ok) {
        const errorMsg = await response.text();
        showPopup(true, errorMsg, 3000);
        throw new Error("Failed to create group");
      }

      // clear form
      setTitle("");
      setDescription("");

      // show successful popup message
      const groupname = formData.get("title");
      showPopup(false, `Group "${groupname}" successfully created.`, 3000);

      // update display
      await fetchMemberships(loggedInUsername);
      await fetchGroups(loggedInUsername);
    } catch (error) {
      console.error("Error submitting form:", error);
    }
  };

  // FUNCTIONS TO CONTROL POPUP MESSAGE -------------------------------------

  const [popupData, setPopupData] = useState({
    show: false,
    isError: false,
    message: "",
    time: 3000,
  });
  function showPopup(isError, message, time = 3000) {
    setPopupData({ show: true, isError, message, time });
  }
  function handlePopupClose() {
    setPopupData((prev) => ({ ...prev, show: false }));
  }

  return (
    <div className="groupMenu-page w-full">
      <div className="flex flex-col items-start w-full">
        {/* Request List Section (if the user has requests) */}
        {requestsLoading ? (
          <p className="mb-2">Loading Group Requests...</p>
        ) : (
          requests &&
          requests.length != 0 && (
            <GroupRequests
              requests={requests}
              onAcceptOrReject={async (request, status) => {
                setRequestsLoading(true);
                if (status == "approve") {
                  showPopup(
                    false,
                    `"${request.username}" is now a member of "${request.title}"!`,
                    3000
                  );
                } else if (status == "reject") {
                  showPopup(
                    true,
                    `Reject request from "${request.username}".`,
                    3000
                  );
                }
                console.log("updating dates...");
                setTimeout(async () => {
                  await fetchGroups(loggedInUsername);
                }, 1000);
              }}
            />
          )
        )}

        {/* Invitation List Section (if the user has invitations) */}
        {membershipsLoading ? (
          <p className="mb-2">Loading Group Invitations...</p>
        ) : (
          invitations &&
          invitations.length != 0 && (
            <GroupInvitation
              invitations={invitations}
              onAcceptOrReject={async (invitationTitle, status) => {
                setMembershipsLoading(true);
                if (status == "approve") {
                  showPopup(
                    false,
                    `You are now a member of "${invitationTitle}"!`,
                    3000
                  );
                } else if (status == "reject") {
                  showPopup(
                    true,
                    `Reject invitation from "${invitationTitle}".`,
                    3000
                  );
                }
                console.log("updating dates...");
                setTimeout(async () => {
                  await fetchMemberships(loggedInUsername);
                  await fetchGroups(loggedInUsername);
                }, 1000);
              }}
            />
          )
        )}

        {/* Create New Group Section */}
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

        {/* User Made Group List Section */}
        <h2 className="mt-8 text-lg text-accent font-bold">Groups You Made</h2>
        {groupsUserMade.length === 0 ? (
          <p>No groups yet</p>
        ) : (
          <RenderList
            items={groupsUserMade.map((group) => group.title)}
            showMoreState={showMoreUserMadeGroups}
            setShowMoreState={setShowMoreUserMadeGroups}
            type="group"
          />
        )}

        {/* User Joined Group List Section */}
        <h2 className="mt-8 text-lg text-accent font-bold">
          Groups You Joined
        </h2>
        {joinedGroups.length === 0 ? (
          <p>No groups yet</p>
        ) : (
          <RenderList
            items={joinedGroups.map((group) => group.title)}
            showMoreState={showMoreUserJoinedGroups}
            setShowMoreState={setShowMoreUserJoinedGroups}
            type="group"
          />
        )}

        {/* All Group List Section */}
        <h2 className="mt-8 text-lg text-accent font-bold">
          Browse All Groups
        </h2>
        {groups.length === 0 ? (
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
      {/* Popup Message */}
      {popupData.show && (
        <Popup
          isError={popupData.isError}
          message={popupData.message}
          time={popupData.time}
          onClose={handlePopupClose}
        />
      )}{" "}
    </div>
  );
}
