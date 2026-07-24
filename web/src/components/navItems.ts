export interface NavChildItem {
  label: string;
  href: string;
  tabId: string;
}

export interface NavItem {
  label: string;
  href: string;
  tabId: string;
  icon: "home" | "tech" | "gaming" | "travelling" | "writing";
  children?: NavChildItem[];
}

export const navItems: NavItem[] = [
  { label: "Home", href: "/", tabId: "home", icon: "home" },
  { label: "Tech", href: "/tech", tabId: "tech", icon: "tech" },
  { label: "Gaming", href: "/gaming", tabId: "gaming", icon: "gaming" },
  { label: "Travelling", href: "/travelling", tabId: "travelling", icon: "travelling" },
  {
    label: "Writing",
    href: "/writing/practice",
    tabId: "writing",
    icon: "writing",
    children: [{ label: "Practice", href: "/writing/practice", tabId: "writing-practice" }],
  },
];
