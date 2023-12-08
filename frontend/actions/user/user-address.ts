import qs from 'query-string';
import useSWR from 'swr';
import { axiosFetcher } from '@/actions/fetcher';
import { Response, Address } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/user-address`;

interface Query {
  userId: string;
}

interface UseGetAllAddressesReturn {
  userAddressList: Response<Address[]> | undefined;
  isError: unknown;
}

export const useGetAllAddresses = (query: Query): UseGetAllAddressesReturn => {
  const url = qs.stringifyUrl({
    url: URL,
    query: {
      userId: query.userId,
    },
  });
  const { data, error } = useSWR<UseGetAllAddressesReturn['userAddressList']>(url, axiosFetcher, {
    suspense: true,
  });

  return {
    userAddressList: data,
    isError: error,
  };
};
