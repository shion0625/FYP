import useSWR from 'swr';
import { axiosPostFetcher } from '@/actions/fetcher';
import { Response } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/sign-in/`;

export interface SignUpBody {
  userName: string;
  firstName: string;
  lastName: string;
  age: number;
  email: string;
  phone: string;
  password: string;
  confirmPassword: string;
}

interface UseSignUpReturn {
  signUp: (body: SignUpBody) => Promise<Response<null>>;
  isError: unknown;
}

export const UseSignUp = (): UseSignUpReturn => {
  const { error, mutate } = useSWR(URL);

  const signUp = async (body: SignUpBody): Promise<Response<null>> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  return {
    signUp,
    isError: error,
  };
};
