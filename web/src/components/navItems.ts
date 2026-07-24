export interface NavItem {
  label: string;
  href: string;
  tabId: string;
  icon: "home" | "tech" | "gaming" | "travelling";
}

export const navItems: NavItem[] = [
  { label: "Home", href: "/", tabId: "home", icon: "home" },
  { label: "Tech", href: "/tech", tabId: "tech", icon: "tech" },
  { label: "Gaming", href: "/gaming", tabId: "gaming", icon: "gaming" },
  { label: "Travelling", href: "/travelling", tabId: "travelling", icon: "travelling" },
];
