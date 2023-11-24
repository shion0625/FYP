import useSWR from "swr";
import { axiosFetcher } from "@/actions/fecher";
import { Response, ProductItem } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/products`;

interface UseGetProductItemsReturn {
  productItems: Response<ProductItem[]> | undefined;
  isError: any;
}

export const useGetProductItems = (id: string): UseGetProductItemsReturn => {
  const { data, error } = useSWR<Response<ProductItem[]>>(
    `${URL}/${id}/items/`,
    axiosFetcher,
    {
      suspense: true,
    }
  );

  return {
    productItems: data,
    isError: error,
  };
};
