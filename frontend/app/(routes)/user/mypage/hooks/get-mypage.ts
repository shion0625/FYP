import useSWR from 'swr';
import { axiosFetcher } from '@/actions/fetcher';
import { Response, Address, User, PaymentMethod } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/mypage`;

interface UseGetMyPageReturn {
  userProfile: Response<User> | undefined;
  userAddressList: Response<Address[]> | undefined;
  userPaymentMethod: Response<PaymentMethod[]> | undefined;
  isError: unknown;
}

export const UseGetMyPage = (): UseGetMyPageReturn => {
  const { data, error } = useSWR<{
    userProfile: UseGetMyPageReturn['userProfile'];
    userAddressList: UseGetMyPageReturn['userAddressList'];
    userPaymentMethod: UseGetMyPageReturn['userPaymentMethod'];
  }>(URL, axiosFetcher, {
    suspense: true,
  });

  return {
    userProfile: data?.userProfile,
    userAddressList: data?.userAddressList,
    userPaymentMethod: data?.userPaymentMethod,
    isError: error,
  };
};
