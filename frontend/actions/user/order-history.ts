import qs from 'query-string';
import useSWR from 'swr';
import { axiosFetcher } from '@/actions/fetcher';
import { Order } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/order/history`;

interface Query {
  pageNumber?: number;
  count?: number;
}

interface UseGetOrderHistoryReturn {
  getUserOrderHistory: (query: Query) => Promise<Order[]>;
  isError: unknown;
}

export const UseGetOrderHistory = (): UseGetOrderHistoryReturn => {
  const { error, mutate } = useSWR(URL);

  const getUserOrderHistory = async (query: Query): Promise<Order[]> => {
    const url = qs.stringifyUrl({
      url: URL,
      query: {
        pageNumber: query.pageNumber,
        count: query.count,
      },
    });
    const response = await axiosFetcher(url);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  return {
    getUserOrderHistory,
    isError: error,
  };
};
