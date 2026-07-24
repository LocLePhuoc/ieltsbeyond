import { createContext, useContext } from "react";

export interface ActiveTabContextValue {
  activeTab: string;
  setActiveTab: (tab: string) => void;
}

export const ActiveTabContext = createContext<ActiveTabContextValue>({
  activeTab: "",
  setActiveTab: () => {},
});

export function useActiveTab(): ActiveTabContextValue {
  return useContext(ActiveTabContext);
}
