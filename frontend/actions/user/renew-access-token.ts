import useSWR from 'swr';
import { axiosPostFetcher } from '@/actions/fetcher';
import { Response, TokenResponse } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/renew-access-token`;

interface UseRenewAccessTokenReturn {
  renewAccessToken: () => Promise<Response<TokenResponse>>;
  isError: unknown;
}

export const UseRenewAccessToken = (): UseRenewAccessTokenReturn => {
  const { error, mutate } = useSWR(URL);

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
