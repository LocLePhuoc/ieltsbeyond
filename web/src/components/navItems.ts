export interface NavChildItem {
  label: string;
  href: string;
  tabId: string;
}

export interface NavItem {
  label: string;
  href: string;
  tabId: string;
  icon: "writing";
  children?: NavChildItem[];
}

export const navItems: NavItem[] = [
  {
    label: "Writing",
    href: "/writing",
    tabId: "writing",
    icon: "writing",
    children: [{ label: "Practice", href: "/writing/practice", tabId: "writing-practice" }],
  },
];
