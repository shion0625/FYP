import qs from 'query-string';
import useSWR from 'swr';
import { axiosFetcher } from '@/actions/fetcher';
import { Response, Address } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/user-address`;

interface Query {
  userId: string;
  pageNumber: number;
  count: number;
}

interface UseGetOrderHistoryReturn {
  userAddressList: Response<Address[]> | undefined;
  isError: unknown;
}

export const useGetOrderHistory = (query: Query): UseGetOrderHistoryReturn => {
  const url = qs.stringifyUrl({
    url: URL,
    query: {
      userId: query.userId,
    },
  });
  const { data, error } = useSWR<UseGetOrderHistoryReturn['userAddressList']>(url, axiosFetcher, {
    suspense: true,
  });

  return {
    userAddressList: data,
    isError: error,
  };
};
