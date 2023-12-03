import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

interface SessionState {
  accessToken: string | null;
  isLoggedIn: boolean;
  setAccessToken: (accessToken: string) => void;
  setIsLogin: (isLogin: boolean) => void;
}

const useSession = create(
  persist<SessionState>(
    (set, get) => ({
      accessToken: null,
      isLoggedIn: false,
      setAccessToken: (accessToken: string) => {
        set({ accessToken: accessToken, isLoggedIn: true });
      },
      setIsLogin: (isLogin: boolean) => {
        set({ isLoggedIn: isLogin });
      },
    }),
    {
      name: "session-storage", // unique name
      storage: createJSONStorage(() => localStorage),
    }
  )
);

export default useSession;
