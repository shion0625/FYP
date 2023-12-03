import useSWR from "swr";
import { axiosPostFetcher } from "@/actions/fetcher";
import { Response, TokenResponse } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/renew-access-token`;

interface UseRenewAccessTokenReturn {
  renewAccessToken: (body: Body) => Promise<Response<TokenResponse>>;
  isError: any;
}

export const UseRenewAccessToken = (): UseRenewAccessTokenReturn => {
  const { data, error, mutate } = useSWR(URL);

  const renewAccessToken = async (): Promise<Response<TokenResponse>> => {
    const response = await axiosPostFetcher(URL);
    mutate(response, false);
    // アクセストークンを取得
    const accessToken = response.headers["access_token"];

    // アクセストークンをlocalStorageに保存
    localStorage.setItem("accessToken", accessToken);
    return response.data;
  };

  return {
    renewAccessToken,
    isError: error,
  };
};
