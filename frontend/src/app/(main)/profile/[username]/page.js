"use client";
import { useEffect, useState } from "react";
import ProfileImage from "@/components/profileImage";
import { Select } from "@headlessui/react";
import FetchFromBackend from "@/lib/fetch";
import formatDate from "@/lib/formatDate";
import fetchCredential from "@/lib/fetchCredential";
import ProfilePrivacyToggle from "@/components/profilePrivacyToggle";
import Link from "next/link";

export default function Profile({ params }) {
  const { username } = params;
  const [loginUsername, setLoginUsername] = useState(null);
  const [userData, setUserData] = useState(null);
  const [posts, setPosts] = useState([]);
  const [followers, setFollowers] = useState([]);
  const [following, setFollowing] = useState([]);
  const [option, setOption] = useState("public");
  const [isOwner, setIsOwner] = useState(false);
  const [userNotFound, setUserNotFound] = useState(false);
  const [isPrivateProfile, setIsPrivateProfile] = useState(false);
  const [loading, setLoading] = useState(true);

  // TODO: check the profile is own (compare login info and the page path)
  // show profile & option for change profile public/private
  // if the profile is not own, check the profile is public/private

  // fetch user information
  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      setIsOwner(false);
      setUserNotFound(false);
      console.log(`starting process for ${username}`);

      try {
        // check the login username
        const storedUsername = localStorage.getItem("user");
        console.log("Loaded username from localStorage:", storedUsername);

        // fetch user information from database
        const userResponse = await FetchFromBackend(`/profile/${username}`, {
          method: "GET",
        });
        // show message when the user is not found
        if (!userResponse.ok) {
          console.log(`Failed to fetch ${username}`);
          setUserNotFound(true);
          setLoading(false);
          return;
        }
        const user = await userResponse.json();
        console.log("User Data:", user);
        console.log("User Data:", user.user.nickname);

        // login user & the user is same (own profile)
        if (storedUsername === user.user.nickname) {
          setIsOwner(true);
        }

        setUserData(user.user);
        setPosts(user.posts);

        // TODO: fetch followers & followings
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, [username]);

  // show loading message while the process is going on
  if (loading) {
    return <div>Loading...</div>;
  }

  // when the user not found
  if (userNotFound) {
    return <div>The user doesn't exist.</div>;
  }

  // for private profile
  if (!isOwner && userData && !userData.public) {
    return <div>This profile is private. TODO: follow request button</div>;
  }

  // show profile (public)
  return (
    <div className="profile-page">
      {isOwner && (
        <ProfilePrivacyToggle
          initialPublicStatus={userData.public}
          username={username}
        />
      )}
      <div className="flex flex-row items-center">
        {
          <ProfileImage
            alt="Profile Image"
            width={100}
            height={100}
            size={40}
            avatar={"http://localhost:8080/avatar/" + userData.userId}
            className={"rounded-full mr-3 w-auto h-16"}
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
          <p className="ml-2 text-gray-600">No posts yet</p>
        ) : (
          <ul>
            {posts.map((post) => (
              <li className="ml-2 text-gray-600" key={post.postId}>
                <Link
                  href={`/post/${post.postId}`}
                  title={post.subject}
                  className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
                >
                  {post.subject}, {formatDate(post.creationDate)}
                </Link>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}
