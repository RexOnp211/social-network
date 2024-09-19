import Link from "next/link";
import { IoHomeOutline } from "react-icons/io5";
import { MdOutlineGroups } from "react-icons/md";
import { IoChatboxOutline } from "react-icons/io5";
import { IoIosNotificationsOutline } from "react-icons/io";
import { CgProfile } from "react-icons/cg";

const links = [
  { name: "Home", href: "/", icon: IoHomeOutline },
  { name: "Clubs", href: "/clubs", icon: MdOutlineGroups },
  { name: "Chat", href: "/messages", icon: IoChatboxOutline },
  {
    name: "Notifications",
    href: "/notifications",
    icon: IoIosNotificationsOutline,
  },
  { name: "Profile", href: "/profile/username", icon: CgProfile },
];

export default function TopBar() {
  return (
    <div className="bg-primary m-3 p-4 rounded-lg shadow-lg">
      <nav className="flex justify-center">
        {links.map((link) => {
          return (
            <Link
              title={link.name}
              href={link.href}
              key={link.name}
              className="text-foreground transition-colors hover:text-accent ease-in mx-5"
            >
              <link.icon size={32} />
            </Link>
          );
        })}
      </nav>
    </div>
  );
}
