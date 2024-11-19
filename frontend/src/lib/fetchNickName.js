import FetchFromBackend from "@/lib/fetch";

export default async function Fetchnickname(userId) {
  const res = await FetchFromBackend(`/user/${userId}`);
  return res;
}
