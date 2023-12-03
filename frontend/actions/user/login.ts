import useSWR from "swr";
import { axiosPostFetcher } from "@/actions/fetcher";
import { Response, TokenResponse } from "@/types";
import useSession from "@/hooks/use-session";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/login/`;

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
  const session = useSession();

  const login = async (body: Body): Promise<Response<TokenResponse>> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false);
    // アクセストークンを取得
    const accessToken = response.headers["access_token"];
    session.setAccessToken(accessToken);
    return response.data;
  };

  return {
    login,
    isError: error,
  };
};
