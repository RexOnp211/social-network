"use client";

import TopBar from "@/components/topbar";
import SideBar from "@/components/sidebar";
import CreatePost from "@/components/createPost";
import FetchFromBackend from "@/lib/fetch";
import Image from "next/image";
import { useEffect, useState } from "react";
import ProfileImage from "@/components/profileImage";

export default function Home() {
  const [post, setPost] = useState(0);
  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const res = await FetchFromBackend("/");
        const jsonData = await res.json();
        console.log(jsonData);
        setPost(jsonData);
      } catch (error) {
        console.error(error);
      }
    };
    fetchPosts();
  }, []);
  return (
    <>
      <TopBar />
      <div className="flex w-auto">
        <SideBar />
        <div className="m-3 w-[90vw] h-[87vh] text-txtColor bg-primary rounded-lg shadow-lg p-6 overflow-y-auto">
          <h1>Home Page </h1>
          <CreatePost />
          {post.length > 0 ? (
            post.map((post) => (
              <div
                key={post.postId}
                className="bg-secondary p-4 rounded-lg m-4"
              >
                <a
                  className="flex flex-row items-center"
                  href={`/profile/${post.userId}`}
                >
                  <ProfileImage
                    alt={post.subject}
                    width={100}
                    height={100}
                    size={40}
                    userId={post.userId}
                  />
                  {post.userId}
                </a>
                <h1 className="text-xl font-bold">{post.subject}</h1>
                <p>{post.content}</p>
                {post.image ? (
                  <Image
                    src={"http://localhost:8080" + post.image}
                    alt="post image"
                    width={500}
                    height={500}
                    className="w-auto h-48"
                  />
                ) : (
                  ""
                )}
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
