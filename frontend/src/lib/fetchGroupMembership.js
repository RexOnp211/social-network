import FetchFromBackend from "@/lib/fetch";

export default async function FetchGroupMembership(
  loggedInUsername,
  decodedGroupname,
  groupOwner
) {
  console.log("memberShip", loggedInUsername, decodedGroupname, groupOwner);

  // show owner content login user == group creator
  if (loggedInUsername === groupOwner) {
    return "owner";
  }

  const memberStatResponse = await FetchFromBackend(
    `/fetch_memberships/${loggedInUsername}`,
    {
      method: "GET",
      credentials: "include"
    }
  );
  if (!memberStatResponse.ok) {
    throw new Error(`Failed to fetch membership ${loggedInUsername}`);
  }
  const memberStat = await memberStatResponse.json();
  console.log("memberStat", memberStat.memberships);

  // get member status for this group
  const groupMembership = memberStat.memberships.find((membership) => {
    return membership.title === decodedGroupname;
  });
  console.log(groupMembership);

  if (groupMembership) {
    console.log("Found membership for group:", groupMembership);

    switch (groupMembership.status) {
      case "approved":
        console.log("User is an approved member of the group.");
        return "approved";
      case "invited":
        console.log("User has been invited to the group.");
        return "invited";
      case "requested":
        console.log("User has requested to join the group.");
        return "requested";
      default:
        console.error("User has an unknown status.");
        return "none";
    }
  }
}
