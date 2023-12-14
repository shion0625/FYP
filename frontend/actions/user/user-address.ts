import * as yup from 'yup';
import useSWR from 'swr';
import { axiosFetcher, axiosPostFetcher, axiosPutFetcher } from '@/actions/fetcher';
import { Address } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/user/address`;

export const AddressSchema = yup.object().shape({
  name: yup
    .string()
    .required('Name is required')
    .min(2, 'Name must be at least 2 characters')
    .max(100, "Name can't be longer than 100 characters"),
  phoneNumber: yup
    .string()
    .required('Phone number is required')
    .matches(/^\+[1-9]\d{1,14}$/, 'Must be a valid E.164 format for phone number'),
  house: yup.string().required('House is required'),
  area: yup.string().required('area is required'),
  landMark: yup.string().required('Landmark is required'),
  city: yup.string().required('city is required'),
  pincode: yup.string().required('Pincode is required'),
  countryName: yup.string().required('Country name is required'),
});

export const UpdateAddressSchema = AddressSchema.concat(
  yup.object({
    id: yup.number().required('ID is required'),
  })
);

export interface AddressBody extends yup.InferType<typeof AddressSchema> {}

export interface UpdateAddressBody extends yup.InferType<typeof UpdateAddressSchema> {}

interface UseUserAddressesReturn {
  getUserAddress: () => Promise<Address[]>;
  saveUserAddress: (body: AddressBody) => Promise<null>;
  updateUserAddress: (body: UpdateAddressBody) => Promise<null>;
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

  const updateUserAddress = async (body: UpdateAddressBody): Promise<null> => {
    const response = await axiosPutFetcher(URL, body);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  return {
    getUserAddress,
    saveUserAddress,
    updateUserAddress,
    isError: error,
  };
};
