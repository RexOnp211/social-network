"use client";

import { useState, useEffect } from "react";
import FetchFromBackend from "@/lib/fetch";

export default function CreateGroupPost({
  loggedInUsername,
  loggedInUserID,
  groupTitle,
  onPostSubmit,
  showPopup,
}) {
  const [postTitle, setPostTitle] = useState("");
  const [postBody, setPostBody] = useState("");

  async function handleSubmit(e) {
    e.preventDefault();

    try {
      const form = new FormData(e.target);
      form.append("groupname", groupTitle);
      form.append("user_id", loggedInUserID);
      form.append("nickname", loggedInUsername);

      console.log([...form.entries()]);

      FetchFromBackend("/create_group_post", {
        method: "POST",
        body: form,
      });
    } catch (error) {
      console.error(error);
    }

    // clear form
    setPostTitle("");
    setPostBody("");

    showPopup(false, "Post created successfully!", 3000);

    // update posts
    if (onPostSubmit) {
      setTimeout(async () => {
        await onPostSubmit(groupTitle);
      }, 1000);
    }
  }

  return (
    <div className="mt-4 mb-4">
      <div
        className={
          "transition-max-height duration-500 ease-in-out overflow-hidden max-h-screen w-full"
        }
      >
        <form
          id="CreateGroupPost"
          className="flex flex-col justify-center border border-gray-300 p-6 rounded-lg shadow-md bg-white"
          onSubmit={handleSubmit}
          encType="multipart/form-data"
        >
          {/* Post Title */}
          <label
            htmlFor="postTitle"
            className="mb-2 text-lg font-semibold text-gray-700"
          >
            Post Title <span className="text-red-500">*</span>
          </label>
          <input
            type="text"
            id="postTitle"
            name="postTitle"
            placeholder="Enter post title"
            value={postTitle}
            onChange={(e) => setPostTitle(e.target.value)}
            required
            className="mb-4 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-accent focus:bg-primary"
          />

          {/* Post Content */}
          <label
            htmlFor="postBody"
            className="mb-2 text-lg font-semibold text-gray-700"
          >
            Post Body <span className="text-red-500">*</span>
          </label>
          <textarea
            id="postBody"
            name="postBody"
            placeholder="Enter post body"
            value={postBody}
            onChange={(e) => setPostBody(e.target.value)}
            required
            className="mb-4 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-accent focus:bg-primary"
          ></textarea>

          {/* Image Uploader */}
          <label
            htmlFor="image"
            className="mb-2 text-lg font-semibold text-gray-700"
          >
            Upload Image
          </label>
          <input
            type="file"
            id="image"
            name="image"
            accept=".png, .jpg, .jpeg, .gif"
            className="mb-4 w-full p-2 border border-gray-300 rounded-md focus:bg-primary"
          />

          <button
            type="submit"
            className="bg-accent w-full text-white rounded-lg p-3 transition-colors hover:bg-accentDark"
          >
            Create Post
          </button>
        </form>
      </div>
    </div>
  );
}
