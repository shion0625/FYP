import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';

interface LoginStore {
  isLogin: boolean;
  onLogin: () => void;
  onLogout: () => void;
}

const useLoginState = create(
  persist<LoginStore>(
    (set) => ({
      isLogin: false,
      onLogin: () => set({ isLogin: true }),
      onLogout: () => set({ isLogin: false }),
    }),
    {
      name: 'login-storage',
      storage: createJSONStorage(() => localStorage),
    }
  )
);

export default useLoginState;
