"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import FetchGroupInfo from "@/lib/fetchGroupInfo";
import FetchGroupMembership from "@/lib/fetchGroupMembership";
import FetchGroupPosts from "@/lib/fetchGroupPosts";
import RenderPost from "@/lib/renderPost";
import GroupPost from "@/components/groupPost";
import Popup from "@/components/popup";

export default function Group({ params }) {
  const { groupname, postId } = params;
  const router = useRouter();
  const decodedGroupname = decodeURIComponent(groupname);
  const [loggedInUsername, setLoggedInUsername] = useState(
    localStorage.getItem("user")
  );
  const [groupData, setGroupData] = useState(null);
  const [posts, setPosts] = useState(null);
  const [memberStatus, setMemberStatus] = useState("none");

  useEffect(() => {
    async function loadData() {
      // fetch group information from database
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
        updatePosts(decodedGroupname);
      } else {
        router.replace(`/group/${groupname}`);
        return;
      }
    }
    loadData();
  }, []);

  // function to update group posts
  async function updatePosts(groupname) {
    try {
      const fetchedPosts = await FetchGroupPosts(groupname);
      setPosts(fetchedPosts);
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
    <div className="groupPost-page">
      {/* Main Section */}
      <div className="flex flex-row items-center">
        <div className="ml-2">
          {memberStatus == "none" ? (
            <div>Loading...</div>
          ) : (
            <>
              {/* Back to Group */}
              <Link
                href={`/group/${decodedGroupname}`}
                title={`${decodedGroupname}`}
                className="text-foreground transition-colors hover:text-accent ease-in hover:underline"
              >
                ◀︎ Back to the group
              </Link>

              <h1 className="mt-2 text-2xl text-accent font-bold">
                {decodedGroupname} Group Posts
              </h1>

              {/* Group Post Form */}
              {
                <GroupPost
                  loggedInUsername={loggedInUsername}
                  groupTitle={decodedGroupname}
                  onPostSubmit={updatePosts}
                  showPopup={showPopup}
                />
              }

              {/* Group Posts */}
              {posts &&
                posts.length > 0 &&
                posts.map((post) => (
                  <div key={post.Id}> {RenderPost(post, groupData)}</div>
                ))}
            </>
          )}
        </div>
        {/* Popup Message */}
        {popupData.show && (
          <Popup
            isError={popupData.isError}
            message={popupData.message}
            time={popupData.time}
            onClose={handlePopupClose}
          />
        )}{" "}
      </div>
    </div>
  );
}
