"use client";
import { useRouter } from "next/navigation";
import FetchGroupInfo from "@/lib/fetchGroupInfo";
import FetchGroupMembership from "@/lib/fetchGroupMembership";
import { FetchGroupPost, FetchGroupPostComments } from "@/lib/fetchGroupPosts";
import { useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import ProfileImage from "@/components/profileImage";
import CreateGroupPostComment from "@/components/createGroupPostComment";
import Popup from "@/components/popup";
import UploadImage from "@/components/images";

export default function GroupPost({ params }) {
  const { groupname, postID } = params;
  const router = useRouter();
  const decodedGroupname = decodeURIComponent(groupname);
  const [groupData, setGroupData] = useState(null);
  const [memberStatus, setMemberStatus] = useState("none");
  const [post, setPost] = useState();
  const [comments, setComments] = useState({});
  const [loggedInUserID, setLoggedInUserID] = useState(
    localStorage.getItem("userID")
  );
  const [loggedInUsername, setLoggedInUsername] = useState(
    localStorage.getItem("user")
  );
  // fetch post
  useEffect(() => {
    async function loadData() {
      console.log(params);
      console.log(postID);

      // fetch group info
      const groupInfo = await FetchGroupInfo(decodedGroupname);
      console.log(groupInfo);
      setGroupData(groupInfo);

      // fetch login user's member status
      const membership = await FetchGroupMembership(
        loggedInUsername,
        decodedGroupname,
        groupInfo.creatorName
      );
      console.log(membership);
      setMemberStatus(membership);
      if (membership == "owner" || membership == "approved") {
        const groupData = await FetchGroupPost(postID);
        console.log(groupData);
        setPost(groupData);
        setComments(await FetchGroupPostComments(postID));
        return;
      }

      // no membership
      router.push(`/group/${groupname}`);
    }
    loadData();
  }, [decodedGroupname, groupname, loggedInUsername, params, postID, router]);

  // function to update comments
  async function updateComments(postID) {
    try {
      const fetchedComments = await FetchGroupPostComments(postID);
      setComments(fetchedComments);
    } catch (error) {
      console.error("Failed to fetch group posts:", error);
    }
  }

  // FUNCTIONS TO CONTROL POPUP MESSAGE -------------------------------------

  const [popupData, setPopupData] = useState({
    show: false,
    isError: false,
    message: "",
    time: 3000,
  });
  function showPopup(isError, message, time = 3000) {
    setPopupData({ show: true, isError, message, time });
  }
  function handlePopupClose() {
    setPopupData((prev) => ({ ...prev, show: false }));
  }

  return (
    <>
      {/* Back Links */}
      <Link
        href={`/group/${decodedGroupname}`}
        title={`${decodedGroupname}`}
        className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
      >
        ◀︎ Back to the group
      </Link>
      <br></br>
      <Link
        href={`/group/${decodedGroupname}/group-post`}
        title={`${decodedGroupname}`}
        className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
      >
        ◀︎ Back to the group posts
      </Link>
      {/* Post */}
      {post && (
        <div key={post.postId} className="bg-secondary p-4 rounded-lg m-4">
          <a
            className="flex flex-row items-center"
            href={`/profile/${post.nickname}`}
          >
            <ProfileImage
              alt={post.subject}
              width={100}
              height={100}
              size={40}
              avatar={"/avatar/" + post.userId}
              className={"rounded-full mr-3 w-auto h-16"}
            />
            {post.nickname || "loading..."}
          </a>
          <h1 className="text-xl font-bold">{post.subject}</h1>
          <p>{post.content}</p>
          {post.image ? (
            <UploadImage
              upload={"/group-post-image/" + post.image}
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
      <CreateGroupPostComment
        postId={postID}
        loggedInUserID={loggedInUserID}
        loggedInUsername={loggedInUsername}
        onCommentSubmit={updateComments}
        showPopup={showPopup}
      />
      {/* Comments */}
      {comments &&
        comments.length > 0 &&
        comments.map((comment) => (
          <div
            key={comment.commentId}
            className="bg-secondary p-2 rounded-lg m-2"
          >
            <a
              className="flex flex-row items-center"
              href={`/profile/${comment.nickname}`}
            >
              <ProfileImage
                alt={comment.subject}
                width={100}
                height={100}
                size={40}
                avatar={"/avatar/" + comment.userId}
                className={"rounded-full mr-3 w-auto h-16"}
              />
              {comment.nickname || "loading..."}
            </a>
            <p>{comment.content}</p>
            {comment.image ? (
              <UploadImage
                upload={"/group-post-comment-image/" + comment.image}
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
      {/* Popup Message */}
      {popupData.show && (
        <Popup
          isError={popupData.isError}
          message={popupData.message}
          time={popupData.time}
          onClose={handlePopupClose}
        />
      )}{" "}
    </>
  );
}
