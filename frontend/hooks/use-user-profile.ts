import { create } from 'zustand';
import { User, Address } from '@/types';

interface UserProfileStore {
  user?: User;
  addressList: Address[];
  setUser: (data: User) => void;
  setAddressList: (data: Address[]) => void;
}

const useUserProfile = create<UserProfileStore>((set) => ({
  user: undefined,
  setUser: (data: User) => set({ user: data }),
  addressList: [],
  setAddressList: (data: Address[]) => set({ addressList: data }),
}));

export default useUserProfile;
