import { Link } from "react-router-dom";
import { navItems } from "./navItems";
import NavIcon from "./NavIcon";

export default function Sidebar({ activeTab }: { activeTab: string }) {
  return (
    <aside className="hidden md:flex flex-col fixed left-0 top-0 h-screen w-[264px] p-3 z-40">
      <SidebarInner activeTab={activeTab} />
    </aside>
  );
}

export function SidebarInner({ activeTab }: { activeTab: string }) {
  return (
    <div className="flex flex-col h-full bg-white/80 backdrop-blur-xl rounded-[28px] shadow-glass px-5 py-6 border border-white/60 relative overflow-hidden">
      {/* Decorative gradient blobs */}
      <div className="absolute -top-20 -right-20 w-40 h-40 bg-gradient-to-br from-sage-200/30 to-navy-200/20 rounded-full blur-3xl pointer-events-none" />
      <div className="absolute -bottom-16 -left-16 w-32 h-32 bg-gradient-to-tr from-navy-200/20 to-sage-100/20 rounded-full blur-3xl pointer-events-none" />

      {/* Profile section */}
      <div className="flex flex-col items-center mb-6 pb-5 border-b border-gray-100/80 relative z-10">
        <div className="relative mb-3">
          <div className="w-14 h-14 rounded-2xl bg-gradient-to-br from-navy-400 to-sage rotate-3 flex items-center justify-center text-white text-lg font-serif font-bold shadow-soft-lg transition-transform duration-300 hover:rotate-6">
            B
          </div>
          <div className="absolute -bottom-1 -right-1 w-4 h-4 bg-sage-400 rounded-full border-2 border-white" />
        </div>
        <span className="font-sans text-sm font-bold text-charcoal tracking-tight">My Blog</span>
        <span className="text-[11px] text-charcoal-light/60 mt-0.5 tracking-wide">THOUGHTS & STORIES</span>
      </div>

      {/* Navigation */}
      <nav className="flex flex-col gap-0.5 relative z-10">
        {navItems.map((item) => {
          const active = activeTab === item.tabId;
          return (
            <Link
              key={item.tabId}
              to={item.href}
              className={
                active
                  ? "sidebar-link flex items-center gap-3 px-3.5 py-2.5 rounded-xl text-[13px] font-semibold text-sage bg-sage-50/80 transition-all duration-200"
                  : "sidebar-link flex items-center gap-3 px-3.5 py-2.5 rounded-xl text-[13px] font-medium text-charcoal-light hover:text-charcoal hover:bg-gray-50/80 transition-all duration-200"
              }
            >
              <NavIcon icon={item.icon} />
              {item.label}
              {active && (
                <div className="ml-auto w-1.5 h-1.5 rounded-full bg-sage ring-2 ring-sage-200/60" />
              )}
            </Link>
          );
        })}
      </nav>

      {/* Bottom */}
      <div className="mt-auto pt-4 relative z-10">
        <div className="flex items-center justify-center gap-3 text-charcoal-light/30">
          <div className="h-px flex-1 bg-gradient-to-r from-transparent via-gray-200 to-transparent" />
          <span className="text-[10px] font-medium tracking-widest uppercase">Blog</span>
          <div className="h-px flex-1 bg-gradient-to-r from-transparent via-gray-200 to-transparent" />
        </div>
      </div>
    </div>
  );
}
