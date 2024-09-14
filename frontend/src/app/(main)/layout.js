import SideBar from "@/components/sidebar";
import TopBar from "@/components/topbar";

export default function DefaultLayout({ children }) {
  return (
    <div>
      <TopBar />
      <SideBar />
      <div>{children}</div>
    </div>
  );
}
