"use client";
import { useEffect, useState } from "react";
import FetchFromBackend from "@/lib/fetch";
import Link from "next/link";

export default function Group({ params }) {
  const { groupname } = params;
  const decodedGroupname = decodeURIComponent(groupname);
  const [groupData, setGroupData] = useState(null);
  const [invited, setInvited] = useState("");
  const [isOwner, setIsOwner] = useState(false);
  const [memberStatus, setMemberStatus] = useState("none");
  const [loading, setLoading] = useState(true);
  const [popupMessage, setPopupMessage] = useState("");
  const [showPopup, setShowPopup] = useState(false);
  const [isError, setIsError] = useState(false);
  const [loggedInUsername, setloggedInUsername] = useState(
    localStorage.getItem("user")
  );

  // fetch group information
  useEffect(() => {
    async function loadData() {
      setLoading(true);
      setIsOwner(false);

      console.log(`starting process for ${decodedGroupname}`);

      // fetch group information from groups table
      const response = await FetchFromBackend(`/group/${decodedGroupname}`, {
        method: "GET",
      });
      if (!response.ok) {
        throw new Error(`Failed to fetch ${decodedGroupname}`);
      }
      const data = await response.json();
      console.log(data.group);
      setGroupData(data.group);

      // login user == group creator
      if (loggedInUsername === data.group.creatorName) {
        setIsOwner(true);
      }

      // check login user status for the group
      const memberStatResponse = await FetchFromBackend(
        `/group_member/${loggedInUsername}`,
        {
          method: "GET",
        }
      );
      if (!memberStatResponse.ok) {
        console.error(
          "Failed to fetch group_member, response status:",
          memberStatResponse.status
        );
        throw new Error(`Failed to fetch group_member`);
      }
      const memberStat = await memberStatResponse.json();
      console.log("memberStat", memberStat.groupmembers);

      const groupMembership = memberStat.groupmembers.find((member) => {
        console.log(member.title, ":", decodedGroupname);
        return member.title === decodedGroupname;
      });
      console.log(groupMembership);

      setMemberStatus("none");
      if (groupMembership) {
        console.log("Found membership for group:", groupMembership);

        switch (groupMembership.Status) {
          case "approved":
            console.log("User is an approved member of the group.");
            setMemberStatus("approved");
            break;
          case "invited":
            console.log("User has been invited to the group.");
            setMemberStatus("invited");
            break;
          case "requested":
            console.log("User has requested to join the group.");
            setMemberStatus("requested");
            break;
          default:
            console.log("User has an unknown status.");
        }
      }

      setLoading(false);
    }
    loadData();
  }, [decodedGroupname, memberStatus]);

  // Request to join group
  const requestToJoin = async () => {
    changeMemberStatus(decodedGroupname, loggedInUsername, "requested");
  };

  // submit invitation
  const OnSubmit = async (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    // stop submitting to user-self
    if (loggedInUsername === formData.get("invited")) {
      setPopupMessage("You cannot invite yourself to the group.");
      setIsError(true);
      setShowPopup(true);

      // Hide the message after 3 seconds
      setTimeout(() => {
        setShowPopup(false);
      }, 5000);

      return;
    }

    changeMemberStatus(decodedGroupname, formData.get("invited"), "invited");
  };

  async function changeMemberStatus(groupname, username, status) {
    try {
      const response = await FetchFromBackend("/invitemember", {
        method: "POST",
        body: JSON.stringify({
          groupname: groupname,
          username: username,
          status: status,
        }),
      });
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "An error occurred");
      }

      // show successful popup message
      setPopupMessage(`Invitation successfully sent to ${invited}.`);
      setIsError(false);
      setShowPopup(true);
    } catch (error) {
      // show error message popup
      setPopupMessage(error.message);
      setIsError(true);
      setShowPopup(true);
    }

    // Hide the message after 3 seconds
    setTimeout(() => {
      setShowPopup(false);
    }, 5000);
  }

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
          <h1 className="text-2xl text-accent font-bold">{groupData.title}</h1>
          <p className="text-gray-600">{groupData.description}</p>
          {isOwner ? (
            <form
              onSubmit={OnSubmit}
              className="mt-4 w-full bg-orange-100 p-4 rounded-lg"
            >
              <div className="w-full">
                <label
                  htmlFor="invited"
                  className="mb-2 text-lg text-accent font-bold"
                >
                  Invite Users to Group
                </label>
                <input
                  type="text"
                  id="invited"
                  name="invited"
                  value={invited}
                  onChange={(e) => setInvited(e.target.value)}
                  required
                  className="p-2 rounded-lg w-full max-w-full"
                  placeholder="Enter username"
                />
              </div>
              <button
                type="submit"
                className="mt-2 transition-colors ease-in hover:bg-accentDark bg-accent text-white rounded-lg p-2 w-full"
              >
                Send Invitation
              </button>
            </form>
          ) : (
            <div className="mt-4 w-full bg-orange-100 p-4 rounded-lg">
              {memberStatus === "approved" && <p>TODO: show posts & events</p>}
              {memberStatus === "invited" && (
                <Link
                  href={"/group"}
                  title="You are invited to the group. Respond to your group invitation on your group menu."
                  className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
                >
                  You are invited to the group. Respond to your group invitation
                  on your group menu.
                </Link>
              )}
              {memberStatus === "requested" && (
                <p>
                  You sent join request to the group. Wait for the owner accept
                  your request.
                </p>
              )}
              {memberStatus === "none" || memberStatus === undefined ? (
                <button
                  onClick={requestToJoin}
                  className="mt-2 transition-colors ease-in hover:bg-accentDark bg-accent text-white rounded-lg p-2"
                >
                  Sent Join Request to Group
                </button>
              ) : null}
            </div>
          )}
        </div>
      </div>

      {showPopup && (
        <div
          className={`absolute bottom-0 left-0 p-4 m-10 rounded shadow-lg transition-all ${
            isError ? "bg-red-500 text-white" : "bg-green-500 text-black"
          }`}
        >
          {popupMessage}
        </div>
      )}
    </div>
  );
}
