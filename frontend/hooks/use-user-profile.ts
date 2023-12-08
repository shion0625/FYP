import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';

interface UserIdStore {
  userId: string;
  setUserId: (data: string) => void;
}

const useUserId = create(
  persist<UserIdStore>(
    (set) => ({
      userId: '',
      setUserId: (userId: string) => set({ userId }),
    }),
    {
      name: 'user-storage',
      storage: createJSONStorage(() => localStorage),
    }
  )
);

export default useUserId;
