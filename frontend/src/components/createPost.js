"use client";

import { useState, useEffect } from "react";
import { Select } from "@headlessui/react";
import FetchFromBackend from "@/lib/fetch";
import fetchCredential from "@/lib/fetchCredential";

export default function CreatePost() {
  const [option, setOption] = useState("public");
  const [followers, setFollowers] = useState([])

  function OnChange(e) {
    setOption(e.target.value);
  }

  async function handleSubmit(e) {
    e.preventDefault();
    try {
      const form = new FormData(e.target);
      console.log("formcreatepost",)
      FetchFromBackend("/", {
        method: "POST",
        credentials: "include",
        body: form,
      })
      .then(() => window.location.reload())
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    const checkFollowers = async () => {
      const user = await fetchCredential()

      const followerRes = await FetchFromBackend(`/followers/${user.username}`, {
        method: "GET",
        credentials: "include"
      })
      const followersJson = await followerRes.json()
      setFollowers(followersJson)
      console.log("users current followers", followers, user.username)
    }
    checkFollowers()
  }, [])

  return (
    <div className="mb-4">
      <div
        className={
          "transition-max-height duration-500 ease-in-out overflow-hidden max-h-screen w-full"
        }
      >
        <form
          id="CreatePost"
          className="flex flex-col justify-center border border-gray-300 p-6 rounded-lg shadow-md bg-white"
          onSubmit={handleSubmit}
        >
          <label
            htmlFor="postTitle"
            className="mb-2 text-lg font-semibold text-gray-700"
          >
            Post Title
          </label>
          <input
            type="text"
            id="postTitle"
            name="postTitle"
            placeholder="Enter post title"
            className="mb-4 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-accent focus:bg-primary"
          />

          <label
            htmlFor="postBody"
            className="mb-2 text-lg font-semibold text-gray-700"
          >
            Post Body
          </label>
          <textarea
            id="postBody"
            name="postBody"
            placeholder="Enter post body"
            className="mb-4 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-accent focus:bg-primary"
          ></textarea>

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
            className="mb-4 w-full p-2 border border-gray-300 rounded-md focus:bg-primary"
          />
          <label
            htmlFor="privacy"
            className="mb-2 text-lg font-semibold text-gray-700"
          >
            Privacy
          </label>
          <Select
            name="privacy"
            className="mb-4 p-2 border border-gray-300 rounded-md focus:bg-primary focus:outline-none focus:ring-2 focus:ring-accent"
            aria-label="profile-type"
            value={option}
            onChange={OnChange}
          >
            <option value="public">Public</option>
            <option value="almost private">almost private</option>
            <option value="private">private</option>
          </Select>
          {option === "private" ? 
          <div>
            <p>Choose which followers see this</p>
            {followers.map((follower) => (
            <div key={follower.id}>
            <input type="checkbox" name="followers" value={follower.id} key={follower.id}/>
            {" " + follower.nickname}
            </div>
            )
            )}
          </div>
          : ""}
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
