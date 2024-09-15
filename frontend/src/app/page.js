import TopBar from "@/components/topbar";
import SideBar from "@/components/sidebar";
import FetchFromBackend from "@/lib/fetch";

export default async function Home() {
  const res = await FetchFromBackend("/");
  return (
    <>
      <TopBar />
      <SideBar />
      <h1>Home Page </h1>
      <p>{res.message}</p>
    </>
  );
}
