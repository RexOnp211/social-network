import FetchFromBackend from "@/lib/fetch";

export default async function FetchGroupPosts(groupname) {
  console.log("fetching posts");
  const postResponse = await FetchFromBackend(
    `/fetch_group_posts/${groupname}`,
    {
      method: "GET",
      credentials: "include"
    }
  );
  if (!postResponse.ok) {
    throw new Error(`Failed to fetch group posts ${groupname}`);
  }
  const posts = await postResponse.json();
  console.log("posts", posts);
  return posts.groupPosts;
}

export async function FetchGroupPostComments(postID) {
  console.log("comments");
  const commentsResponse = await FetchFromBackend(
    `/fetch_group_post_comment/${postID}`,
    {
      method: "GET",
      credentials: "include"
    }
  );
  if (!commentsResponse.ok) {
    throw new Error(`Failed to fetch group post comments ${postID}`);
  }
  const comments = await commentsResponse.json();
  console.log("comments", comments);
  return comments.comments;
}

export async function FetchGroupPost(postID) {
  console.log("fetching post");
  const postResponse = await FetchFromBackend(`/fetch_group_post/${postID}`, {
    method: "GET",
  });
  if (!postResponse.ok) {
    throw new Error(`Failed to fetch group post ${postID}`);
  }
  const post = await postResponse.json();
  console.log("post", post);
  return post.post;
}
