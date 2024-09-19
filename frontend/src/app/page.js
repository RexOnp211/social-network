import TopBar from "@/components/topbar";
import SideBar from "@/components/sidebar";
import FetchFromBackend from "@/lib/fetch";

export default async function Home() {
  const res = await FetchFromBackend("/");
  return (
    <>
      <TopBar />
      <div className="flex w-auto">
        <SideBar />
        <div className="m-3 w-[90vw] text-txtColor bg-primary rounded-lg shadow-lg p-6">
          <h1>Home Page </h1>
          <p>{res.message}</p>
        </div>
      </div>
    </>
  );
}
