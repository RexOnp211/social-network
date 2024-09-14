import Link from "next/link";

const links = [
  { name: "Home", href: "/" },
  { name: "Clubs", href: "/clubs" },
  { name: "Chat", href: "/messages" },
  { name: "Notifications", href: "/notifications" },
];

export default function TopBar() {
  return (
    <div>
      {links.map((link) => {
        return (
          <Link href={link.href} key={link.name}>
            {link.name}
          </Link>
        );
      })}
    </div>
  );
}
