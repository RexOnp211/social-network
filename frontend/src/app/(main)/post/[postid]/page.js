"use client";
import FetchFromBackend from "@/lib/fetch";
import { useEffect, useState } from "react";
import Image from "next/image";
import ProfileImage from "@/components/profileImage";
import CreateComment from "@/components/createComment";
import Fetchnickname from "@/lib/fetchNickName";

export default function Posts({ params }) {
  const [post, setPost] = useState({});
  const [nickname, setNickname] = useState("");
  const [commentNN, setCommentNN] = useState({});
  useEffect(() => {
    const fetchPost = async () => {
      try {
        const path = `/post/${params.postid}`;
        const res = await FetchFromBackend(path, {
          method: "GET",
        });
        const jsonData = await res.json();
        console.log(jsonData);
        setPost(jsonData);
      } catch (error) {
        console.error(error);
      }
    };
    fetchPost();
  }, [params.postid]);

  useEffect(() => {
    const findNickname = async () => {
      const id = post.userId;
      try {
        const res = await Fetchnickname(id);
        const textData = await res.text();
        setNickname(textData);
      } catch (error) {
        console.error(`Error fetching nickname for userId ${id}:`, error);
      }
    };

    if (post.userId) {
      findNickname();
    }
  }, [post]);

  useEffect(() => {
    const findNickname = async () => {
      const comments = post.comments || [];

      const ids = [...new Set(comments.map((p) => p.userId))];
      console.log(ids);
      const commentNNmap = {};

      for (const userId of ids) {
        console.log("userid in comments", userId);
        try {
          const res = await Fetchnickname(userId);
          const textData = await res.text();
          commentNNmap[userId] = textData;
          console.log("this is textdata from commetns ids", commentNNmap);
        } catch (error) {
          console.error(`Error fetching nickname for userId ${userId}:`, error);
        }
      }

      setCommentNN(commentNNmap);
    };
    const com = post.comments || [];
    if (com.length > 0) {
      findNickname();
    }
  }, [post.comments]);

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
              avatar={`${process.env.NEXT_PUBLIC_API_URL}/avatar/${post.userId}`}
              className={"rounded-full mr-3 w-auto h-16"}
            />
            {nickname || "loading..."}
          </a>
          <h1 className="text-xl font-bold">{post.subject}</h1>
          <p>{post.content}</p>
          {post.image ? (
            <Image
              src={`${process.env.NEXT_PUBLIC_API_URL}/image/ + post.image`}
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
      {post &&
        post.comments &&
        post.comments.map((comment) => (
          <div
            key={comment.commentId}
            className="bg-secondary p-2 rounded-lg m-2"
          >
            <a
              className="flex flex-row items-center"
              href={`/profile/${comment.userId}`}
            >
              <ProfileImage
                alt={comment.subject}
                width={100}
                height={100}
                size={40}
                avatar={`${process.env.NEXT_PUBLIC_API_URL}/avatar/" + comment.userId`}
                className={"rounded-full mr-3 w-auto h-16"}
              />
              {commentNN[comment.userId] || "loading..."}
            </a>
            <p>{comment.content}</p>
            {comment.image ? (
              <Image
                src={`${process.env.NEXT_PUBLIC_API_URL}/image/" + comment.image`}
                alt="comment image"
                width={500}
                height={500}
                className="w-auto h-48"
              />
            ) : (
              ""
            )}
          </div>
        ))}
    </>
  );
}
