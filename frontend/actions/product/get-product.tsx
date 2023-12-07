import useSWR from 'swr';
import { Response, Product } from '@/types';
import { axiosFetcher } from '@/actions/fetcher';

const URL = `${process.env.NEXT_PUBLIC_API_URL}/products`;

interface UseGetProductReturn {
  product: Response<Product> | undefined;
  isLoading: boolean;
  isError: unknown;
}

export const useGetProduct = (id: string): UseGetProductReturn => {
  const { data, isLoading, error } = useSWR<Response<Product>>(`${URL}/${id}`, axiosFetcher, {
    suspense: true,
  });

  return {
    product: data,
    isLoading: isLoading,
    isError: error,
  };
};
