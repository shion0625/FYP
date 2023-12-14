import useSWR from 'swr';
import { axiosFetcher } from '@/actions/fetcher';
import { Response, User } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/user/profile`;

interface UseGetProfileReturn {
  userProfile: Response<User> | undefined;
  isError: unknown;
}

export const useGetProfile = (): UseGetProfileReturn => {
  const { data, error } = useSWR<UseGetProfileReturn['userProfile']>(URL, axiosFetcher, {
    suspense: true,
  });

  return {
    userProfile: data,
    isError: error,
  };
};
