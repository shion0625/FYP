import useSWR from "swr";
import { Response, Product } from "@/types";
import { axiosFetcher } from "@/actions/fecher";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/products`;

interface UseGetProductReturn {
  product: Response<Product> | undefined;
  isError: any;
}

export const useGetProduct = (id: string): UseGetProductReturn => {
  const { data, error } = useSWR<Response<Product>>(
    `${URL}/${id}`,
    axiosFetcher,
    {
      suspense: true,
    }
  );

  return {
    product: data,
    isError: error,
  };
};
