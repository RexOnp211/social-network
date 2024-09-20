"use client";

import ProfileImage from "@/components/profileImage";
import { Select } from "@headlessui/react";
import { useState } from "react";

export default function Profile({ params }) {
  const alt = "Profile Image";
  const img = "/image/profile-default.png";
  const [option, setOption] = useState("public");
  const OnChange = (e) => {
    setOption(e.target.value);
  };
  return (
    <div className="flex flex-row items-center">
      <ProfileImage src={img} width={500} height={500} alt={alt} size={124} />
      <div>
        <h1 className="ml-4 text-2xl">userName</h1>
        <form className="ml-4">
          <Select
            name="privacy"
            className="border data-[hover]:shadow data-[focus]:bg-accent w-20 h-8"
            aria-label="profile-type"
            value={option}
            onChange={OnChange}
          >
            <option value="public">public</option>
            <option value="private">private</option>
          </Select>
        </form>
      </div>
    </div>
  );
}
