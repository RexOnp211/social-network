"use client";

import TopBar from "@/components/topbar";
import SideBar from "@/components/sidebar";
import CreatePost from "@/components/createPost";
import FetchFromBackend from "@/lib/fetch";
import { useEffect, useState } from "react";

export default function Home() {
  const [post, setPost] = useState(0);
  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const res = await FetchFromBackend("/");
        const jsonData = await res.json();
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
              <div key={post.id} className="bg-secondary p-4 rounded-lg m-4">
                <h1 className="text-xl font-bold">{post.title}</h1>
                <p>{post.postBody}</p>
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
