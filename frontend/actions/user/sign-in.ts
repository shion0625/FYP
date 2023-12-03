import qs from "query-string";
import useSWR from "swr";
import { axiosPostFetcher } from "@/actions/fecher";
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

interface UseSignInReturn {
  signIn: Response<any> | undefined;
  isError: any;
}

export const UseSignIn = (): {
  signIn: (body: Body) => Promise<UseSignInReturn>;
  isError: any;
} => {
  const { data, error, mutate } = useSWR(URL);

  const signIn = async (body: Body) => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response;
  };

  return {
    signIn,
    isError: error,
  };
};