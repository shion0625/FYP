import qs from 'query-string';
import useSWR from 'swr';
import { axiosFetcher } from '@/actions/fetcher';
import { Response, User } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/user-profile`;

interface Query {
  userId: string;
}

interface UseGetProfileReturn {
  userProfile: Response<User> | undefined;
  isError: unknown;
}

export const useGetProfile = (query: Query): UseGetProfileReturn => {
  const url = qs.stringifyUrl({
    url: URL,
    query: {
      userId: query.userId,
    },
  });
  const { data, error } = useSWR<UseGetProfileReturn['userProfile']>(url, axiosFetcher, {
    suspense: true,
  });

  return {
    userProfile: data,
    isError: error,
  };
};
