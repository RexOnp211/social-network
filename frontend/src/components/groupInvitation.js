"use client";
import { useEffect, useState } from "react";
import Popup from "@/components/popup";
import UpdateMembership from "@/lib/updateMembership";
import Link from "next/link";

const GroupInvitation = ({ invitations, onAcceptOrReject }) => {
  const [loggedInUsername, setLoggedInUsername] = useState(
    localStorage.getItem("user")
  );

  async function handleInvitationStatus(invitation, status) {
    try {
      const resultStatus = UpdateMembership(
        invitation.id,
        invitation.title,
        loggedInUsername,
        status
      );
      console.log(invitation);

      onAcceptOrReject(invitation.title, status);
    } catch (error) {
      console.error(`Error updating status ${status}:`, error);
    }
  }

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
    <div className="bg-orange-100 p-4 mb-4 rounded-lg">
      <h2 className="text-lg text-accent font-bold mb-2">
        You have received group invitations:
      </h2>
      <div>
        {invitations.length !== 0 && (
          <ul className="list-disc pl-5 marker:text-txtColor">
            {invitations.map((invitation) => (
              <li className="ml-2 text-gray-600 mb-2" key={invitation.title}>
                <Link
                  href={`/group/${invitation.title}`}
                  className="text-txtColor hover:underline font-bold mr-4"
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
};

export default GroupInvitation;
