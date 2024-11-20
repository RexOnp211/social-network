"use client";
import { useEffect, useState } from "react";
import RenderList from "@/components/renderList";
import FetchFromBackend from "@/lib/fetch";
import Link from "next/link";

const GroupInvitation = ({ invitationExist, username, onAcceptOrReject }) => {
  const [showMoreState, setShowMoreState] = useState(false);
  const [popupMessage, setPopupMessage] = useState("");
  const [showPopup, setShowPopup] = useState(false);

  async function handleInvitationStatus(invitation, status) {
    try {
      console.log(typeof invitation.id);
      const id =
        typeof invitation.id === "string"
          ? Number(invitation.id)
          : invitation.id;

      const response = await FetchFromBackend("/update_member_status", {
        headers: {
          "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify({
          id: id,
          status: status,
        }),
      });
      if (!response.ok) {
        throw new Error(`Failed to update status ${status}`);
      }

      // show successful popup message
      let message = `You are now a member of ${invitation.title}!`;
      if (status == "reject") {
        message = `Reject invitation from ${invitation.title}`;
      }
      setPopupMessage(message);
      setShowPopup(true);

      onAcceptOrReject();

      // Hide the message after 3 seconds
      setTimeout(() => {
        setShowPopup(false);
      }, 3000);
    } catch (error) {
      console.error(`Error updating status ${status}:`, error);
    }
  }

  return (
    <div className="bg-orange-100 p-4 mb-4 rounded-lg">
      <h2 className="text-lg text-accent font-bold mb-2">
        You have received group invitations:
      </h2>

      <div>
        {invitationExist.length !== 0 && (
          <ul className="list-disc pl-5 marker:text-txtColor">
            {invitationExist.map((invitation) => (
              <li className="ml-2 text-gray-600 mb-2" key={invitation.title}>
                <Link
                  href={`/group/${invitation.title}`}
                  className="text-txtColor hover:underline mr-4"
                >
                  {invitation.title}
                </Link>
                <button
                  onClick={() => handleInvitationStatus(invitation, "approve")}
                  className="mr-2 transition-colors ease-in hover:bg-accentDark bg-accent text-white rounded-lg p-2 py-0.5"
                >
                  Accept
                </button>
                <button
                  onClick={() => handleInvitationStatus(invitation, "reject")}
                  className="transition-colors ease-in hover:bg-red-700 bg-red-600 text-white rounded-lg p-2 py-0.5"
                >
                  Reject
                </button>
              </li>
            ))}
          </ul>
        )}
      </div>

      {showPopup && (
        <div className="absolute bottom-0 left-0 p-4 mb-10 ml-10 bg-orange-100 text-black rounded shadow-lg">
          {popupMessage}
        </div>
      )}
    </div>
  );
};

export default GroupInvitation;
