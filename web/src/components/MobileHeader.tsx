import { useState } from "react";
import { Link, useLocation } from "react-router-dom";
import { navItems } from "./navItems";
import NavIcon from "./NavIcon";

export default function MobileHeader({ activeTab }: { activeTab: string }) {
  const [open, setOpen] = useState(false);
  const location = useLocation();

  return (
    <header className="md:hidden fixed top-0 left-0 right-0 z-50 px-3 pt-3">
      <div className="bg-white/80 backdrop-blur-xl rounded-2xl shadow-glass border border-white/60 px-4 py-3">
        <div className="flex items-center justify-between">
          <Link to="/" className="flex items-center gap-2">
            <div className="w-7 h-7 rounded-lg bg-gradient-to-br from-navy-400 to-sage flex items-center justify-center text-white text-[11px] font-serif font-bold">
              B
            </div>
            <span className="font-sans text-sm font-bold text-charcoal">My Blog</span>
          </Link>
          <button
            type="button"
            onClick={() => setOpen((v) => !v)}
            className="text-charcoal-light p-1.5 rounded-xl hover:bg-gray-100/80 transition-colors"
          >
            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth={1.8}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
            </svg>
          </button>
        </div>
        {open && (
          <nav className="mt-3 flex flex-col gap-0.5 border-t border-gray-100/80 pt-3">
            {navItems.map((item) => {
              const active = activeTab === item.tabId;
              return (
                <div key={item.tabId}>
                  <Link
                    to={item.href}
                    onClick={() => setOpen(false)}
                    className={
                      active
                        ? "flex items-center gap-3 px-3.5 py-2.5 rounded-xl text-[13px] font-semibold text-sage bg-sage-50/80"
                        : "flex items-center gap-3 px-3.5 py-2.5 rounded-xl text-[13px] text-charcoal-light hover:bg-gray-50/80 transition-colors"
                    }
                  >
                    <NavIcon icon={item.icon} />
                    {item.label}
                  </Link>
                  {item.children && (
                    <div className="flex flex-col gap-0.5 mt-0.5 ml-[15px] pl-4 border-l border-gray-200/70">
                      {item.children.map((child) => {
                        const childActive = location.pathname === child.href;
                        return (
                          <Link
                            key={child.tabId}
                            to={child.href}
                            onClick={() => setOpen(false)}
                            className={
                              childActive
                                ? "flex items-center px-3 py-2 rounded-xl text-[13px] font-semibold text-sage bg-sage-50/80"
                                : "flex items-center px-3 py-2 rounded-xl text-[13px] text-charcoal-light hover:bg-gray-50/80 transition-colors"
                            }
                          >
                            {child.label}
                          </Link>
                        );
                      })}
                    </div>
                  )}
                </div>
              );
            })}
          </nav>
        )}
      </div>
    </header>
  );
}
