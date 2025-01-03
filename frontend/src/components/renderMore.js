import Link from "next/link";

function RenderMore(posts) {
  const maxPostsToShow = 3;
  const postsToShow = posts.slice(0, maxPostsToShow);

  return (
    <div>
      {/* Render few posts */}
      {postsToShow.map((post) => (
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
              avatar={"http://localhost:8080/avatar/" + post.userId}
              className={"rounded-full mr-3 w-auto h-16"}
            />
            {post.nickname || "loading..."}
          </a>
          <h1 className="text-xl font-bold">{post.subject}</h1>
          <p>{post.content}</p>
          {post.image ? (
            <Image
              src={"http://localhost:8080/image/" + post.image}
              alt="post image"
              width={500}
              height={500}
              className="w-auto h-80"
            />
          ) : null}
          <Link
            href={`${encodeURIComponent(groupTitle)}/group-post/${post.Id}`}
            title="comments"
          >
            <IoChatboxOutline />
          </Link>
        </div>
      ))}

      {/* Show more & create new post */}
      <div className="text-center mt-4">
        <Link
          href={`/group/${encodeURIComponent(groupTitle)}/group-post`}
          className="text-accent underline"
        >
          Show more & Create new post
        </Link>
      </div>
    </div>
  );
}

const RenderList = ({ items, showMoreState, setShowMoreState, type }) => {
  return (
    <>
      <ul className="list-disc pl-5 marker:text-txtColor">
        {(showMoreState ? items : items.slice(0, 4)).map((item, index) => (
          <li key={index}>
            <Link
              href={`/${
                type === "event"
                  ? "events"
                  : type === "group"
                  ? "group"
                  : "profile"
              }/${item}`}
              className="text-txtColor hover:underline"
            >
              {item}
            </Link>
          </li>
        ))}
      </ul>
      {items.length > 4 && showMore(showMoreState, setShowMoreState)}
    </>
  );
};

const showMore = (state, setState) => {
  return (
    <button
      onClick={() => setState(!state)}
      className="hover:underline flex items-center"
    >
      {state ? "Show Less" : "Show More"}
      <span className={`ml-1 transform ${state ? "rotate-180" : ""}`}>▼</span>
    </button>
  );
};

export default RenderList;
