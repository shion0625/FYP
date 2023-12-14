import useSWR from 'swr';
import { axiosFetcher, axiosPostFetcher } from '@/actions/fetcher';
import { Address } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/user/address`;

export interface AddressBody {
  name: string;
  area: string;
  city: string;
  countryName: string;
  house: string;
  landMark: string;
  phoneNumber: string;
  pincode: string;
}

interface UseUserAddressesReturn {
  getUserAddress: () => Promise<Address[]>;
  saveUserAddress: (body: AddressBody) => Promise<null>;
  isError: unknown;
}

export const UseUserAddresses = (): UseUserAddressesReturn => {
  const { error, mutate } = useSWR(URL);

  const saveUserAddress = async (body: AddressBody): Promise<null> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  const getUserAddress = async (): Promise<Address[]> => {
    const response = await axiosFetcher(URL);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  return {
    getUserAddress,
    saveUserAddress,
    isError: error,
  };
};
