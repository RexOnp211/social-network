import React, { useEffect } from "react";

// show popup message for limited time
export default function Popup({ isError, message, onClose, time }) {
  useEffect(() => {
    const timer = setTimeout(() => {
      onClose();
    }, time);

    return () => clearTimeout(timer);
  }, [onClose, time]);

  return (
    <div
      className={`fixed bottom-0 left-0 p-4 m-10 rounded shadow-lg transition-all ${
        isError ? "bg-red-500 text-white" : "bg-green-500 text-black"
      }`}
    >
      {message}
    </div>
  );
}
