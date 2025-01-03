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
import { useRouter } from "next/navigation";


const links = [
  { name: "Home", href: "/", icon: IoHomeOutline },
  { name: "Groups", href: "/group", icon: MdOutlineGroups },
  { name: "Chat", href: "/messages", icon: IoChatboxOutline },
  {
    name: "Notifications",
    href: "/notifications",
    icon: IoIosNotificationsOutline,
  },
  {
    name: "Profile",
    href: (loggedInUsername) => `/profile/${loggedInUsername}`,
    icon: CgProfile,
  },
];

export default function TopBar() {
  // fetch login username and use it for profile link
  const [loggedInUsername, setLoggedInUsername] = useState("");

  useEffect(() => {
    const checklogin = async () => {
      const res = await fetchCredential()
      const data = await res.username
      setLoggedInUsername(data)
    }
    checklogin()

  }, []);

  const router = useRouter();

  const Logout = async () => {
    try {
      const res = await FetchFromBackend("/logout", {
        method: "POST",
        credentials: "include",
      });
      if (res.status === 200) {
        console.log("Logout successful");

        // clear login info from local storage
        localStorage.removeItem("userID");
        localStorage.removeItem("user");

        router.push("/login");
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
            typeof link.href === "function"
              ? link.href(loggedInUsername)
              : link.href;
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
      <button onClick={Logout} className="bg-secondary p-2 rounded-lg">
        Logout
      </button>
    </div>
  );
}
