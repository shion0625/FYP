import * as yup from 'yup';
import useSWR from 'swr';
import { axiosFetcher, axiosPutFetcher } from '@/actions/fetcher';
import { User } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/user/profile`;

export const ProfileSchema = yup.object().shape({
  userName: yup
    .string()
    .required('Username is required')
    .matches(/^[a-zA-Z0-9]*$/, 'Only alphanumeric characters are allowed for username')
    .min(3, 'Username must be at least 3 characters')
    .max(15, "Username can't be longer than 15 characters"),
  firstName: yup
    .string()
    .required('First name is required')
    .matches(/^[a-zA-Z]*$/, 'Only alphabets are allowed for first name')
    .min(2, 'First name must be at least 2 characters')
    .max(50, "First name can't be longer than 50 characters"),
  lastName: yup
    .string()
    .required('Last name is required')
    .matches(/^[a-zA-Z]*$/, 'Only alphabets are allowed for last name')
    .min(1, 'Last name must be at least 1 character')
    .max(50, "Last name can't be longer than 50 characters"),
  age: yup
    .number()
    .required('Age is required')
    .min(0, "Age can't be less than 0")
    .max(120, "Age can't be more than 120")
    .transform((value, originalValue) => {
      return originalValue === '' ? undefined : value;
    }),
  email: yup.string().required('Email is required').email('Must be a valid email address'),
  phone: yup
    .string()
    .required('Phone number is required')
    .matches(/^\+[1-9]\d{1,14}$/, 'Must be a valid E.164 format for phone number'),
});

export interface ProfileBody extends yup.InferType<typeof ProfileSchema> {}

interface UseGetProfileReturn {
  getUserProfile: () => Promise<User>;
  updateUserProfile: (body: ProfileBody) => Promise<null>;
  isError: unknown;
}

export const UseGetProfile = (): UseGetProfileReturn => {
  const { error, mutate } = useSWR(URL);

  const getUserProfile = async (): Promise<User> => {
    const response = await axiosFetcher(URL);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  const updateUserProfile = async (body: ProfileBody): Promise<null> => {
    const response = await axiosPutFetcher(URL, body);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  return {
    getUserProfile,
    updateUserProfile,
    isError: error,
  };
};
