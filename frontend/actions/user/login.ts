import useSWR from "swr";
import { axiosPostFetcher } from "@/actions/fetcher";
import { Response, TokenResponse } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/login/`;

interface Body {
  email: string;
  password: string;
}

interface UseLoginReturn {
  login: (body: Body) => Promise<Response<TokenResponse>>;
  isError: any;
}

export const UseLogin = (): UseLoginReturn => {
  const { data, error, mutate } = useSWR(URL);

  const login = async (body: Body): Promise<Response<any>> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false);
    // アクセストークンを取得
    const accessToken = response.headers["access_token"];

    // アクセストークンをlocalStorageに保存
    localStorage.setItem("accessToken", accessToken);
    return response.data;
  };

  return {
    login,
    isError: error,
  };
};
