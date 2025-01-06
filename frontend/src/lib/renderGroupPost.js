import Image from "next/image";
import Link from "next/link";
import ProfileImage from "@/components/profileImage";
import { IoChatboxOutline } from "react-icons/io5";
import UploadImage from "@/components/images";

export default function RenderGroupPost(post, groupData) {
  console.log(post);
  return (
    <div>
      <div key={post.Id} className="bg-secondary p-4 rounded-lg m-4">
        <a
          className="flex flex-row items-center"
          href={`/profile/${post.nickname}`}
        >
          <ProfileImage
            alt={post.subject}
            width={100}
            height={100}
            size={40}
            avatar={"/avatar/" + post.Id}
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
            className="w-auto h-80"
          />
        ) : null}
        <Link
          href={`/group/${groupData.title}/group-post/${post.Id}`}
          title="comments"
        >
          <IoChatboxOutline />
        </Link>
      </div>
    </div>
  );
}
