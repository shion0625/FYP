import useSWR from 'swr';
import { axiosFetcher, axiosPostFetcher } from '@/actions/fetcher';
import { Response, PaymentMethod } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/paymentMethod`;

export interface PaymentMethodBody {
  creditNumber: string;
  cvv: string;
}

interface UsePaymentMethodReturn {
  getPaymentMethod: () => Promise<Response<PaymentMethod[]>>;
  savePaymentMethod: (body: PaymentMethodBody) => Promise<Response<null>>;
  isError: unknown;
}

export const UsePaymentMethod = (): UsePaymentMethodReturn => {
  const { error, mutate } = useSWR(URL);

  const savePaymentMethod = async (body: PaymentMethodBody): Promise<Response<null>> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  const getPaymentMethod = async () => {
    const response = await axiosFetcher(URL);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  return {
    getPaymentMethod,
    savePaymentMethod,
    isError: error,
  };
};
