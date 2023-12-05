import useSWR from "swr";
import { axiosPostFetcher } from "@/actions/fetcher";
import { Response } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/sign-in/`;

interface Body {
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
  signUp: (body: Body) => Promise<Response<any>>;
  isError: any;
}

export const UseSignUp = (): UseSignUpReturn => {
  const { data, error, mutate } = useSWR(URL);

  const signUp = async (body: Body): Promise<Response<any>> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  return {
    signUp,
    isError: error,
  };
};
