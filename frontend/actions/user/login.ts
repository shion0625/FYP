import useSWR from 'swr';
import { axiosPostFetcher } from '@/actions/fetcher';
import { Response, TokenResponse } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/login`;

export interface LoginBody {
  email: string;
  password: string;
}

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
