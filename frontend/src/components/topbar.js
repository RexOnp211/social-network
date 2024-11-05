"use client";

import Link from "next/link";
import { IoHomeOutline } from "react-icons/io5";
import { MdOutlineGroups } from "react-icons/md";
import { IoChatboxOutline } from "react-icons/io5";
import { IoIosNotificationsOutline } from "react-icons/io";
import { CgProfile } from "react-icons/cg";
import { useEffect, useState } from "react";
import fetchCredential from "@/lib/fetchCredential";
import FetchFromBackend from "@/lib/fetch";

const links = [
  { name: "Home", href: "/", icon: IoHomeOutline },
  { name: "Groups", href: "/groupmenu", icon: MdOutlineGroups },
  { name: "Chat", href: "/messages", icon: IoChatboxOutline },
  {
    name: "Notifications",
    href: "/notifications",
    icon: IoIosNotificationsOutline,
  },
  {
    name: "Profile",
    href: (username) => `/profile/${username}`,
    icon: CgProfile,
  },
];

export default function TopBar() {
  // fetch login username and use it for profile link
  const [username, setUsername] = useState("");
  useEffect(() => {
    const fetchUserCredential = async () => {
      const credential = await fetchCredential();
      if (credential && credential.username) {
        setUsername(credential.username);
      }
    };
    fetchUserCredential();
  }, []);

  const logout = async () => {
    try {
      const res = await FetchFromBackend("/logout", {
        method: "POST",
      });
      if (res.status === 200) {
        console.log("Logout successful");
      } else {
        throw new Error("Logout failed");
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="bg-primary m-3 p-4 rounded-lg shadow-lg">
      <nav className="flex justify-center">
        {links.map((link) => {
          const href =
            typeof link.href === "function" ? link.href(username) : link.href;
          return (
            <Link
              title={link.name}
              href={href}
              key={link.name}
              className="text-foreground transition-colors hover:text-accent ease-in mx-5"
            >
              <link.icon size={32} />
            </Link>
          );
        })}
      </nav>
      <button onClick={logout} className="bg-secondary p-2 rounded-lg">
        Logout
      </button>
    </div>
  );
}
