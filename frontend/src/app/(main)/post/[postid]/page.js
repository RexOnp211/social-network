"use client";
import FetchFromBackend from "@/lib/fetch";
import { useEffect, useState } from "react";
import Image from "next/image";
import ProfileImage from "@/components/profileImage";
import { IoChatboxOutline } from "react-icons/io5";
import Link from "next/link";
import CreatePost from "@/components/createPost";
import CreateComment from "@/components/createComment";

export default function Posts({ params }) {
  const [post, setPost] = useState(null);
  useEffect(() => {
    const fetchPost = async () => {
      try {
        const path = `/post/${params.postid}`;
        const res = await FetchFromBackend(path);
        const jsonData = await res.json();
        console.log(jsonData);
        setPost(jsonData);
      } catch (error) {
        console.error(error);
      }
    };
    fetchPost();
  }, [params.postid]);
  return (
    <>
      {post && (
        <div key={post.postId} className="bg-secondary p-4 rounded-lg m-4">
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
              src={"http://localhost:8080/image/" + post.image}
              alt="post image"
              width={500}
              height={500}
              className="w-auto h-48"
            />
          ) : (
            ""
          )}
        </div>
      )}
      <CreateComment postId={params.postid} />
    </>
  );
}
