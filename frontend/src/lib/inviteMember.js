import FetchFromBackend from "@/lib/fetch";

export default async function InviteMember(groupname, username) {
  console.log(groupname, username);
  try {
    const response = await FetchFromBackend("/invite_member", {
      method: "POST",
      body: JSON.stringify({
        groupname: groupname,
        username: username,
      }),
    });
    console.log("response", response);

    if (!response.ok) {
      const errorData = await response.json();
      const msg = errorData.message || "Server error";
      console.log("msg", msg);
      return msg;
    }
  } catch (error) {
    return "An error occurred";
  }

  return "";
}
