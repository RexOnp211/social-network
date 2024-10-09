"use client";
import { useEffect, useState } from "react";
import ProfileImage from "@/components/profileImage";
import { Select } from "@headlessui/react";
import FetchFromBackend from "@/lib/fetch";
import formatDate from "@/lib/formatDate";

export default function Profile({ params }) {
  const { username } = params;
  const [userData, setUserData] = useState(null);
  const [posts, setPosts] = useState([]);
  const [followers, setFollowers] = useState([]);
  const [following, setFollowing] = useState([]);
  const [option, setOption] = useState("public");
  const [isOwner, setIsOwner] = useState(false);
  const [loading, setLoading] = useState(true);

  // TODO: check the profile is own (compare login info and the page path)
  // show profile & option for change profile public/private
  // if the profile is not own, check the profile is public/private

  // fetch user information
  useEffect(() => {
    async function loadData() {
      setLoading(true);
      console.log(`starting process for ${username}`);

      // fetch user information from users table
      const response = await FetchFromBackend(`/profile/${username}`, {
        method: "GET",
      });
      if (!response.ok) {
        throw new Error(`Failed to fetch ${username}`);
      }
      const data = await response.json();
      console.log(data);
      setUserData(data.user);

      // fetch user activity (user's posts)
      setPosts(data.posts);

      // TODO: fetch followers & followings

      setLoading(false);
    }
    loadData();
  }, [username]);

  // don't show profile (private)
  useEffect(() => {
    if (userData && !userData.public) {
      setUserData(null);
    }
  }, [userData]);

  // show loading message while the process is going on
  if (loading) {
    return <div>Loading...</div>;
  }

  // show message when the user is not found
  if (!userData) {
    return <div>The user doesn't exist or their profile is not public.</div>;
  }

  // show profile (public)
  return (
    <div className="profile-page">
      <div className="flex flex-row items-center">
        {
          <ProfileImage
            alt="Profile Image"
            width={50}
            height={50}
            avatar={userData.avatar ? userData.avatar : "profile-default.png"}
          />
        }
        <div className="ml-10">
          <h1 className="text-2xl">{userData.nickname}</h1>
          <p className="text-gray-600">
            {userData.firstname} {userData.lastname}
          </p>
          <p className="text-gray-600">{userData.email}</p>
          <p className="text-gray-600">
            Date of Birth: {formatDate(userData.dob)}
          </p>
        </div>
      </div>

      <div>
        <h2 className="mt-4 text-lg text-accent font-bold">About Me</h2>
        <p className="ml-2 text-gray-600">{userData.aboutMe}</p>
      </div>

      <div>
        <h2 className="mt-4 text-lg text-accent font-bold">User Activity</h2>
        {posts.length === 0 ? (
          <p>No posts yet</p>
        ) : (
          <ul>
            {posts.map((post) => (
              <li className="ml-2 text-gray-600" key={post.post_id}>
                {post.subject}, {formatDate(post.creationDate)}
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}
