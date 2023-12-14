import * as yup from 'yup';
import useSWR from 'swr';
import { axiosPostFetcher } from '@/actions/fetcher';
import { Response, TokenResponse } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/login`;

export const loginSchema = yup.object().shape({
  email: yup.string().required('Email is required').email('Must be a valid email address'),
  password: yup
    .string()
    .required('Password is required')
    .min(5, 'Password must be at least 5 characters')
    .max(30, "Password can't be longer than 30 characters"),
});

export interface LoginBody extends yup.InferType<typeof loginSchema> {}

interface UseLoginReturn {
  login: (body: LoginBody) => Promise<Response<TokenResponse>>;
  isError: unknown;
}

export const UseLogin = (): UseLoginReturn => {
  const { error, mutate } = useSWR(URL);
  const login = async (body: LoginBody): Promise<Response<TokenResponse>> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false);
    return response.data;
  };

  return {
    login,
    isError: error,
  };
};
