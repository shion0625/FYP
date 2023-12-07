import useSWR from "swr";
import { axiosPostFetcher } from "@/actions/fetcher";
import { Response, ProductVariationValue, ProductItemInfo } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/purchase`;

export interface Body {
  userId: string;
  addressId: number;
  productItemInfo: ProductItemInfo[];
  totalFee: number;
  paymentMethodID: number;
}

interface UsePurchaseReturn {
  purchaseOrder: (body: Body) => Promise<Response<unknown>>;
  isError: any;
}

export const UsePurchase = (): UsePurchaseReturn => {
  const { data, error, mutate } = useSWR(URL);

  const purchaseOrder = async (body: Body): Promise<Response<unknown>> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false);
    return response.data;
  };

  return {
    purchaseOrder,
    isError: error,
  };
};
