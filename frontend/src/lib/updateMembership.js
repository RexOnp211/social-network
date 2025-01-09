import FetchFromBackend from "@/lib/fetch";

// helper function to send request to server to change group membership status
export default async function UpdateMembership(
  id,
  groupname,
  username,
  status,
  chatId
) {
  console.log("updating membership", id, groupname, username, status);
  try {
    const response = await FetchFromBackend("/update_membership", {
      method: "POST",
      body: JSON.stringify({
        id: Number(id),
        groupname: groupname,
        username: username,
        status: status,
        chatId: Number(chatId)
      }),
    });
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || "An error occurred");
    }
  } catch (error) {
    return error;
  }

  return "success";
}
