import { useState } from "react";
import FetchFromBackend from "@/lib/fetch";

const ProfilePrivacyToggle = ({ initialPublicStatus, username }) => {
  const [isPublic, setIsPublic] = useState(initialPublicStatus);
  const [popupMessage, setPopupMessage] = useState("");
  const [showPopup, setShowPopup] = useState(false);

  const handleToggle = async () => {
    const newPublicStatus = !isPublic;

    console.log("wasPublic", isPublic);
    console.log("newPublicStatus", newPublicStatus);

    try {
      const formData = new FormData();
      formData.append("username", username);
      formData.append("privacy", newPublicStatus);

      const response = await FetchFromBackend(`/privacy`, {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        console.error("Failed to update profile privacy.");
        return;
      }

      setIsPublic(newPublicStatus);

      // Set the popup message
      const message = newPublicStatus
        ? "Now Anyone can see your profile."
        : "Now only your followers can see your profile.";
      setPopupMessage(message);
      setShowPopup(true);

      // Hide the message after 3 seconds
      setTimeout(() => {
        setShowPopup(false);
      }, 3000);
    } catch (error) {
      console.error("Error updating profile privacy:", error);
    }
  };

  return (
    <div className="bg-orange-100 p-4 mb-4 rounded-lg">
      <div className="flex items-center">
        <span className="mr-2">Make your profile public: </span>

        <div
          onClick={handleToggle}
          className={`relative inline-flex items-center cursor-pointer w-14 h-8 rounded-full transition-colors ${
            isPublic ? "bg-blue-500" : "bg-gray-300"
          }`}
        >
          <span
            className={`inline-block w-6 h-6 bg-white rounded-full transform transition-transform duration-300 ${
              isPublic ? "translate-x-6" : "translate-x-1"
            }`}
          />
        </div>

        {showPopup && (
          <div className="absolute bottom-0 left-0 p-4 mb-10 ml-10 bg-orange-100 text-black rounded shadow-lg">
            {popupMessage}
          </div>
        )}
      </div>
    </div>
  );
};

export default ProfilePrivacyToggle;
