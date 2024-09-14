import TopBar from "@/components/topbar";
import SideBar from "@/components/sidebar";

export default async function Home() {
  const res = await fetch("http://localhost:8080/");
  const posts = await res.json();
  return (
    <>
      <TopBar />
      <SideBar />
      <h1>Home Page </h1>
      <p>{posts.message}</p>
    </>
  );
}
