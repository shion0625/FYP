import useSWR from 'swr';
import { axiosFetcher } from '@/actions/fetcher';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/logout`;

interface UseGetProfileReturn {
  logout: () => Promise<null>;
  isError: unknown;
}

export const useLogout = (): UseGetProfileReturn => {
  const { error, mutate } = useSWR(URL);

  const logout = async (): Promise<null> => {
    const response = await axiosFetcher(URL);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return null;
  };

  return {
    logout,
    isError: error,
  };
};
