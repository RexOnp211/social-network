"use client";

import TopBar from "@/components/topbar";
import SideBar from "@/components/sidebar";
import CreatePost from "@/components/createPost";
import FetchFromBackend from "@/lib/fetch";
import Image from "next/image";
import { useEffect, useState } from "react";
import ProfileImage from "@/components/profileImage";
import { IoChatboxOutline } from "react-icons/io5";
import Link from "next/link";
import { useRouter } from "next/navigation";
import FetchCredential from "@/lib/fetchCredential";
import Fetchnickname from "@/lib/fetchNickName";
import WsClient from "@/lib/wsClient";
import { useRef } from "react";
import UploadImage from "@/components/images";

export default function Home() {
  const [post, setPost] = useState(0);
  const [nickname, setNickname] = useState({});
  const router = useRouter();
  const ws = useRef(null);
  console.log("ENV", process.env.go_api)

  useEffect(() => {
    const load = async () => {
      const wsClient = await WsClient();
      ws.current = wsClient;

      ws.current.onmessage = (event) => {
        const eventData = event.data
        const parsedData = JSON.parse(eventData)
        if (parsedData.type === "group_invite") {
          alert(`New invite to group, go to groups page to see it`)
        }
      };
    };
    load();
  }, []);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const res = await FetchFromBackend("/", {
          method: "GET",
          credentials: "include"
        });
        const jsonData = await res.json();
        console.log("posts", jsonData);
        setPost(jsonData);
      } catch (error) {
        console.error(error);
      }
    };
    fetchPosts();
  }, []);

  useEffect(() => {
    const findNickname = async () => {
      const ids = [...new Set(post.map((p) => p.userId))];
      const nicknameMap = {};

      for (const userId of ids) {
        try {
          const res = await Fetchnickname(userId);
          const textData = await res.text();
          nicknameMap[userId] = textData;
        } catch (error) {
          console.error(`Error fetching nickname for userId ${userId}:`, error);
        }
      }

      setNickname(nicknameMap);
    };
    if (post.length > 0) {
      findNickname();
    }
  }, [post]);

  useEffect(() => {
    const checkLogin = async () => {
      try {
        const res = await FetchCredential();
        if (res.username === "") {
          router.push("/login");
        } else {
          // set login info from local storage
          localStorage.setItem("userID", await res.id);
          localStorage.setItem("user", await res.username);
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
        <div className="m-3 w-[90vw] h-[87vh] text-txtColor bg-primary rounded-lg shadow-lg p-6 overflow-y-auto">
          <h1>Home Page </h1>
          <CreatePost type="Post" />
          {post.length > 0 ? (
            post.map((post) => (
              <div
                key={post.postId}
                className="bg-secondary p-4 rounded-lg m-4"
              >
                <a
                  className="flex flex-row items-center"
                  href={`/profile/${nickname[post.userId]}`}
                >
                  <ProfileImage
                    alt={post.subject}
                    width={100}
                    height={100}
                    size={40}
                    avatar={"/avatar/" + post.userId}
                    className={"rounded-full mr-3 w-auto h-16"}
                  />
                  {nickname[post.userId] || "loading..."}
                </a>
                <h1 className="text-xl font-bold">{post.subject}</h1>
                <p>{post.content}</p>
                {post.image ? (
                  <UploadImage
                    upload={"/image/" + post.image}
                    alt="post image"
                    width={500}
                    height={500}
                    className="w-auto h-80"
                  />

                ) : (
                  ""
                )}
                <Link href={`/post/${post.postId}`} title="comments">
                  <IoChatboxOutline />
                </Link>
              </div>
            ))
          ) : (
            <p>No posts available</p>
          )}
        </div>
      </div>
    </>
  );
}
