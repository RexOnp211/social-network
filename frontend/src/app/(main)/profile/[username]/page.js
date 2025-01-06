"use client";
import { useEffect, useState } from "react";
import ProfileImage from "@/components/profileImage";
import { Select } from "@headlessui/react";
import FetchFromBackend from "@/lib/fetch";
import formatDate from "@/lib/formatDate";
import ProfilePrivacyToggle from "@/components/profilePrivacyToggle";
import { FollowRequest } from "@/lib/wsClient";
import { useRef } from "react";
import WsClient from "@/lib/wsClient";
import Fetchnickname from "@/lib/fetchNickName";
import Link from "next/link";
import fetchCredential from "@/lib/fetchCredential";
import { useRouter } from "next/navigation";

export default function Profile({ params }) {
  const { username } = params;
  // const [loggedInUsername, setLoggedInUsername] = useState(
  //   localStorage.getItem("user")
  // );
  // doesnt seem to work for me :/
  const [loggedInUsername, setLoggedInUsername] = useState(null);
  const [userData, setUserData] = useState(null);
  const [posts, setPosts] = useState([]);
  const [followers, setFollowers] = useState([]);
  const [following, setFollowing] = useState([]);
  const [option, setOption] = useState("public");
  const [isOwner, setIsOwner] = useState(false);
  const [userNotFound, setUserNotFound] = useState(false);
  const [isPrivateProfile, setIsPrivateProfile] = useState(false);
  const [loading, setLoading] = useState(true);
  const [followsUser, setFollowsUser] = useState(false);
  const ws = useRef(null);
  const router = useRouter();

  // TODO: check the profile is own (compare login info and the page path)
  // show profile & option for change profile public/private
  // if the profile is not own, check the profile is public/private

  const sendFollowRequest = async () => {
    if (ws.current) {
      console.log("sending follow request");
      const user = await fetchCredential();
      const id = await user.id; //int
      const followingId = await userData.id; //int
      const publicProfile = userData.public;
      id === followingId
        ? alert("You can't follow yourself")
        : ws.current.send(
            JSON.stringify({
              type: "follow_request",
              payload: new FollowRequest(
                "" + user.id,
                "" + followingId,
                publicProfile
              ), //user.id is converted to string
            })
          );
      alert("sending follow request");
      window.location.reload();
    }
  };

  const unFollowUser = async () => {
    const user = await fetchCredential();
    const followingId = await user.id;
    const data = { follower_id: followingId, followee_id: userData.id };
    const res = await FetchFromBackend("/unfollow", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    console.log("responce", res);
    if (!res.ok) {
      alert("unfollowing User failed");
    } else {
      alert("Successfully unfollowed user");
      window.location.reload();
    }
  };

  // set up websocket
  useEffect(() => {
    const load = async () => {
      const wsClient = await WsClient();
      ws.current = wsClient;

      ws.current.onmessage = (event) => {
        if (event.type === "follow_request") {
          alert(`Message received: ${event.data}`);
        }
      };
    };
    load();
  }, []);

  // fetch user information
  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      setIsOwner(false);
      setUserNotFound(false);
      console.log(`starting process for ${username}`);

      try {
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

        const loggedUser = await fetchCredential();
        // login user & the user is same (own profile)
        if (loggedUser.username === user.user.nickname) {
          setIsOwner(true);
        }

        setUserData(user.user);
        setPosts(user.posts);

        const following = await FetchFromBackend(
          `/following/${user.user.nickname}`,
          {
            method: "GET",
            credentials: "include",
          }
        );

        const followingUsers = await following.json();

        setFollowing(followingUsers);

        const followers = await FetchFromBackend(
          `/followers/${user.user.nickname}`,
          {
            method: "GET",
            credentials: "include",
          }
        );

        const followerUsers = await followers.json();

        setFollowers(followerUsers);

        if (followerUsers.some((obj) => obj.nickname === loggedUser.username)) {
          setFollowsUser(true);
        }
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
    return <div>The user doesnt exist.</div>;
  }

  // for private profile
  if (!isOwner && userData && !userData.public && !followsUser) {
    return <div>This profile is private. <button className="bg-secondary p-2 rounded-lg" onClick={sendFollowRequest}>Follow user to see info</button></div>;
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
            avatar={"/avatar/" + userData.id}
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
        {isOwner ? (
          ""
        ) : followsUser ? (
          <button onClick={unFollowUser}>UnFollow</button>
        ) : (
          <button onClick={sendFollowRequest}>Follow</button>
        )}
      </div>
      <div>
        <h2 className="mt-4 text-lg text-accent font-bold">About Me</h2>
        <p className="ml-2 text-gray-600">{userData.aboutMe}</p>
      </div>
      <div>
        <h2 className="mt-4 text-lg text-accent font-bold">Posts</h2>
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
        <h2 className="mt-4 text-lg text-accent font-bold">Following</h2>
        {following.length === 0 ? (
          <p className="ml-2 text-gray-600">You are not following anybody</p>
        ) : (
          <ul>
            {following.map((user) => (
              <li className="ml-2 text-gray-600" key={user.id}>
                <Link
                  href={`/profile/${user.nickname}`}
                  title={user.nickname}
                  className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
                >
                  {user.nickname}
                </Link>
              </li>
            ))}
          </ul>
        )}
        <h2 className="mt-4 text-lg text-accent font-bold">Followers</h2>
        {followers.length === 0 ? (
          <p className="ml-2 text-gray-600">You dont have any followers</p>
        ) : (
          <ul>
            {followers.map((user) => (
              <li className="ml-2 text-gray-600" key={user.id}>
                <Link
                  href={`/profile/${user.nickname}`}
                  title={user.nickname}
                  className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
                >
                  {user.nickname}
                </Link>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}
