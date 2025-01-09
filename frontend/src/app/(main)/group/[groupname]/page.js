"use client";
import { useEffect, useState } from "react";
import FetchGroupEvents from "@/lib/fetchGroupEvents";
import FetchGroupInfo from "@/lib/fetchGroupInfo";
import FetchGroupMembership from "@/lib/fetchGroupMembership";
import FetchGroupPosts from "@/lib/fetchGroupPosts";
import UpdateMembership from "@/lib/updateMembership";
import InviteMember from "@/lib/inviteMember";
import RenderGroupPost from "@/lib/renderGroupPost";
import RenderEvent from "@/lib/renderEvent";
import Link from "next/link";
import Popup from "@/components/popup";

export default function Group({ params }) {
  const { groupname } = params;
  const decodedGroupname = decodeURIComponent(groupname); // % -> whitespace
  const [groupData, setGroupData] = useState(null);
  const [memberStatus, setMemberStatus] = useState("none");
  const [invitedUser, setInvitedUser] = useState("");
  const [posts, setPosts] = useState(null);
  const [events, setEvents] = useState(null);
  const [loggedInUsername, setLoggedInUsername] = useState(
    localStorage.getItem("user")
  );
  const [loading, setLoading] = useState(true);

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

  // FUNCTIONS TO DISPLAY GROUP INFO -------------------------------------

  // fetch group information
  useEffect(() => {
    async function loadData() {
      console.log(
        `starting process for ${decodedGroupname}, ${loggedInUsername}`
      );

      // fetch group information from database
      const groupInfo = await FetchGroupInfo(decodedGroupname);
      console.log("GROUP INFO", groupInfo);
      setGroupData(groupInfo);

      // fetch login user's member status
      const membership = await FetchGroupMembership(
        loggedInUsername,
        decodedGroupname,
        groupInfo.creatorName,
        groupInfo.chatId
      );
      console.log(membership);
      if (membership == "owner" || membership == "approved") {
        setPosts(await FetchGroupPosts(decodedGroupname));
        setEvents(await FetchGroupEvents(decodedGroupname));
      }

      setMemberStatus(membership);
      setLoading(false);
    }
    loadData();
  }, [decodedGroupname, memberStatus, loggedInUsername]);

  //  when the group is not found
  if (loading) {
    return <div>Loading group information...</div>;
  } else if (!groupData) {
    return <div>The group "{decodedGroupname}" doesn't exist.</div>;
  }

  // FUNCTIONS TO TRIGGER ACTIONS ON THE PAGE -------------------------------------

  // sent request to join this group
  async function requestToJoin() {
    UpdateMembership(0, decodedGroupname, loggedInUsername, "requested", groupData.chatId);
    setMemberStatus("requested");
    showPopup(
      false,
      `Request successfully sent to "${decodedGroupname}".`,
      5000
    );

    
  }

  // send invitation to other user
  async function OnSubmit(e) {
    e.preventDefault();

    const formData = new FormData(e.target);
    const invitedUser = formData.get("invitedUser");

    // stop submitting to user-self
    if (loggedInUsername === invitedUser) {
      showPopup(true, "You cannot invite yourself.", 5000);
      return;
    }

    console.log(groupData);
    // stop inviting the owner
    if (groupData.creatorName === invitedUser) {
      showPopup(true, "You cannot invite the owner of the group.", 5000);
      return;
    }

    const msg = await InviteMember(decodedGroupname, invitedUser);
    console.log("msg", msg);
    if (!msg) {
      showPopup(
        false,
        `Invitation successfully sent to "${invitedUser}".`,
        5000
      );
      return;
    }
    showPopup(true, msg, 5000);
  }

  return (
    <div className="group-page">
      {/* Main Section */}
      <div className="flex flex-row items-center">
        <div className="ml-2">
          {/* Public Section */}
          <h1 className="text-2xl text-accent font-bold">{groupData.title}</h1>
          <p className="mb-4 text-gray-600">{groupData.description}</p>

          {memberStatus == "none" ? (
            <div>Loading...</div>
          ) : (
            <>
              {/* Private Section (Owner & Member only ) */}
              {memberStatus === "approved" || memberStatus === "owner" ? (
                <>
                  {/* Send Invitation */}
                  <form onSubmit={OnSubmit} className="attentionBG">
                    <div className="w-full">
                      <label
                        htmlFor="invitedUser"
                        className="mb-2 text-lg text-accent font-bold"
                      >
                        Invite Users to Group
                      </label>
                      <input
                        type="text"
                        id="invitedUser"
                        name="invitedUser"
                        value={invitedUser}
                        onChange={(e) => setInvitedUser(e.target.value)}
                        required
                        className="p-2 rounded-lg w-full max-w-full"
                        placeholder="Enter username"
                      />
                    </div>
                    <button type="submit" className="basicButton">
                      Send Invitation
                    </button>
                  </form>

                  {/* Group Posts */}
                  <h2 className="mt-4 mb-2 text-lg text-accent font-bold">
                    Group Posts
                  </h2>
                  {posts &&
                    posts.length > 0 &&
                    posts
                      .slice(0, 2)
                      .map((post) => (
                        <div key={post.Id}>
                          {" "}
                          {RenderGroupPost(post, groupData)}
                        </div>
                      ))}
                  <Link
                    href={`/group/${encodeURIComponent(
                      groupData.title
                    )}/group-post`}
                    className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
                  >
                    {posts.length <= 2
                      ? "Create new post"
                      : "Show more & create new post"}
                  </Link>

                  {/* Group Events */}
                  <h2 className="mt-4 mb-2 text-lg text-accent font-bold">
                    Group Events
                  </h2>
                  {events &&
                    events.length > 0 &&
                    events
                      .slice(0, 2)
                      .map((event) => (
                        <RenderEvent key={event.Id} event={event} />
                      ))}
                  <Link
                    href={`/group/${encodeURIComponent(
                      groupData.title
                    )}/group-event`}
                    className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
                  >
                    {events && events.length > 2
                      ? "Show more & create new event"
                      : "Create new event"}
                  </Link>
                </>
              ) : null}

              {/* Join Request Section (Non-member only) */}
              {memberStatus === "none" || memberStatus === undefined ? (
                <button
                  onClick={requestToJoin}
                  className="mt-2 transition-colors ease-in hover:bg-accentDark bg-accent text-white rounded-lg p-2"
                >
                  Sent Join Request to Group
                </button>
              ) : null}

              {/* Invited Member Section (Invited Member only ) */}
              {memberStatus === "invited" ? (
                <div className="attentionBG">
                  <Link
                    href={"/notification"}
                    title="invited"
                    className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
                  >
                    You are invited to this group. Respond to your group
                    invitation on notification page.
                  </Link>
                </div>
              ) : null}

              {/* Requested Member Section (Requested Member only ) */}
              {memberStatus === "requested" ? (
                <div className="attentionBG">
                  You sent join request to this group. Wait for the owner accept
                  your request.
                </div>
              ) : null}
            </>
          )}
        </div>
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
