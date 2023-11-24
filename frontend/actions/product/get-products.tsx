import qs from "query-string";
import useSWR from "swr";
import { axiosFetcher } from "@/actions/fecher";
import { Response, Product } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/products/`;

interface Query {
  categoryId?: string;
  brandId?: string;
}

interface UseGetProductsReturn {
  products: Response<Product[]> | undefined;
  isError: any;
}

export const useGetProducts = (query: Query): UseGetProductsReturn => {
  const url = qs.stringifyUrl({
    url: URL,
    query: {
      category_id: query.categoryId,
      brand_id: query.brandId,
    },
  });

  const { data, error } = useSWR<UseGetProductsReturn["products"]>(
    url,
    axiosFetcher,
    {
      suspense: true,
    }
  );

  return {
    products: data,
    isError: error,
  };
};
