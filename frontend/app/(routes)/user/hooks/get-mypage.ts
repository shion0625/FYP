import useSWR from 'swr';
import { axiosFetcher } from '@/actions/fetcher';
import { Address, User, PaymentMethod } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/user`;

interface Profile {
  userProfile: User | undefined;
  userAddressList: Address[] | undefined;
  userPaymentMethod: PaymentMethod[] | undefined;
}

interface UseGetMyPageReturn {
  getProfile: () => Promise<Profile>;
  isError: unknown;
}

export const UseGetMyPage = (): UseGetMyPageReturn => {
  const { error, mutate } = useSWR(URL);

  const getProfile = async (): Promise<Profile> => {
    const response = await axiosFetcher(URL);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return {
      userProfile: response.userProfile.data,
      userAddressList: response.userAddressList.data,
      userPaymentMethod: response.userPaymentMethod.data,
    };
  };

  return {
    getProfile,
    isError: error,
  };
};
