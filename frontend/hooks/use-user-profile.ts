import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';

interface UserProfileStore {
  userId: string;
  setUserId: (data: string) => void;
}

const useUserProfile = create(
  persist<UserProfileStore>(
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

export default useUserProfile;
