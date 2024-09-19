import SideBar from "@/components/sidebar";
import TopBar from "@/components/topbar";

export default function DefaultLayout({ children }) {
  return (
    <>
      <TopBar />
      <div className="flex w-auto">
        <SideBar />
        <div className="m-3 w-[90vw] bg-primary rounded-lg shadow-lg p-6">
          <div>{children}</div>
        </div>
      </div>
    </>
  );
}
