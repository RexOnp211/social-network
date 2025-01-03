import FetchFromBackend from "@/lib/fetch";

export default async function FetchGroupInfo(decodedGroupname) {
  const groupInfoResponse = await FetchFromBackend(
    `/group/${decodedGroupname}`,
    {
      method: "GET",
    }
  );
  if (!groupInfoResponse.ok) {
    console.error(`Failed to fetch ${decodedGroupname}`);
  }
  const groupInfo = await groupInfoResponse.json();
  console.log(groupInfo.group);
  return groupInfo.group;
}
