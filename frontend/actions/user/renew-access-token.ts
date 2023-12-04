import useSWR from "swr";
import { axiosPostFetcher } from "@/actions/fetcher";
import { Response, TokenResponse } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/user/renew-access-token`;

interface UseRenewAccessTokenReturn {
  renewAccessToken: () => Promise<Response<TokenResponse>>;
  isError: any;
}

export const UseRenewAccessToken = (): UseRenewAccessTokenReturn => {
  const { data, error, mutate } = useSWR(URL);

  const renewAccessToken = async (): Promise<Response<TokenResponse>> => {
    const response = await axiosPostFetcher(URL);
    mutate(response, false);
    return response.data;
  };

  return {
    renewAccessToken,
    isError: error,
  };
};
