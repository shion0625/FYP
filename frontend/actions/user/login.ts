import useSWR from "swr";
import { axiosPostFetcher } from "@/actions/fetcher";
import { Response, TokenResponse } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/user/login`;

export interface Body {
  email: string;
  password: string;
}

interface UseLoginReturn {
  login: (body: Body) => Promise<Response<TokenResponse>>;
  isError: any;
}

export const UseLogin = (): UseLoginReturn => {
  const { data, error, mutate } = useSWR(URL);

  const login = async (body: Body): Promise<Response<TokenResponse>> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false);
    return response.data;
  };

  return {
    login,
    isError: error,
  };
};
