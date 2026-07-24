import { useEffect, useState } from "react";
import { Outlet, useLocation } from "react-router-dom";
import MobileHeader from "./MobileHeader";
import Sidebar from "./Sidebar";
import Footer from "./Footer";
import { ActiveTabContext } from "./ActiveTabContext";

const CATEGORY_TABS = ["tech", "gaming", "travelling"];

function deriveTabFromPath(pathname: string): string {
  if (pathname === "/") return "home";
  const segment = pathname.split("/")[1];
  return CATEGORY_TABS.includes(segment) ? segment : "";
}

export default function Layout() {
  const location = useLocation();
  const [activeTab, setActiveTab] = useState(() => deriveTabFromPath(location.pathname));

  useEffect(() => {
    window.scrollTo({ top: 0 });
    // Post detail pages set their own active tab once the post's category loads.
    if (!location.pathname.startsWith("/posts/")) {
      setActiveTab(deriveTabFromPath(location.pathname));
    }
  }, [location.pathname]);

  return (
    <ActiveTabContext.Provider value={{ activeTab, setActiveTab }}>
      <div className="bg-surface font-sans text-charcoal min-h-screen relative">
        {/* Decorative background blobs */}
        <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
          <div className="absolute top-[-10%] right-[-5%] w-[600px] h-[600px] bg-gradient-to-br from-sage-100/40 to-transparent rounded-full blur-[120px]" />
          <div className="absolute bottom-[-10%] left-[10%] w-[500px] h-[500px] bg-gradient-to-tr from-navy-100/30 to-transparent rounded-full blur-[120px]" />
        </div>

        <MobileHeader activeTab={activeTab} />
        <div className="flex min-h-screen relative z-10">
          <Sidebar activeTab={activeTab} />
          <main id="content" className="flex-1 md:ml-[264px] px-5 md:px-8 lg:px-12 py-6 md:py-10 mt-14 md:mt-0 max-w-6xl">
            <Outlet />
            <Footer />
          </main>
        </div>
      </div>
    </ActiveTabContext.Provider>
  );
}
