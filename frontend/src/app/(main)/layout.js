"use client";

import SideBar from "@/components/sidebar";
import TopBar from "@/components/topbar";
import FetchCredential from "@/lib/fetchCredential";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function DefaultLayout({ children }) {
  const router = useRouter();
  useEffect(() => {
    const checkLogin = async () => {
      try {
        const res = await FetchCredential();
        if (res.username === "") {
          router.push("/login");
        }
      } catch (error) {
        console.error("error checking login", error);
      }
    };
    checkLogin();
  }, [router]);

  return (
    <>
      <TopBar />
      <div className="flex w-auto">
        <SideBar />
        <div className="m-3 w-[90vw] bg-primary rounded-lg shadow-lg p-6">
          <div>{children}</div>
        </div>
      </div>
    </>
  );
}
