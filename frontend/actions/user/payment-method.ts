import useSWR from 'swr';
import { axiosFetcher, axiosPostFetcher } from '@/actions/fetcher';
import { PaymentMethod } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/paymentMethod`;

export interface PaymentMethodBody {
  number: string;
  name: string;
  expiry: string;
  cvc: string;
}

interface UsePaymentMethodReturn {
  getPaymentMethod: () => Promise<PaymentMethod[]>;
  savePaymentMethod: (body: PaymentMethodBody) => Promise<null>;
  isError: unknown;
}

export const UsePaymentMethod = (): UsePaymentMethodReturn => {
  const { error, mutate } = useSWR(URL);

  const savePaymentMethod = async (body: PaymentMethodBody): Promise<null> => {
    const response = await axiosPostFetcher(URL, body);
    mutate(response, false); // Update the local data immediately, but disable revalidation
    return response.data;
  };

  const getPaymentMethod = async (): Promise<PaymentMethod[]> => {
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
