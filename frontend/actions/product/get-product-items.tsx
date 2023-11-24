"use client";
import useSWR from "swr";
import { axiosFetcher } from "@/actions/fecher";
import { ProductItem, Response } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/products`;

const useProductItems = (id: string) => {
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

export default useProductItems;
