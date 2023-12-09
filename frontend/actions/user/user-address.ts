import useSWR from 'swr';
import { axiosFetcher } from '@/actions/fetcher';
import { Response, Address } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/user-address`;

interface UseGetAllAddressesReturn {
  userAddressList: Response<Address[]> | undefined;
  isError: unknown;
}

export const useGetAllAddresses = (): UseGetAllAddressesReturn => {
  const { data, error } = useSWR<UseGetAllAddressesReturn['userAddressList']>(URL, axiosFetcher, {
    suspense: true,
  });

  return {
    userAddressList: data,
    isError: error,
  };
};
