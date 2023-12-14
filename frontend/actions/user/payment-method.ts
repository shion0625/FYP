import * as yup from 'yup';
import useSWR from 'swr';
import { axiosFetcher, axiosPostFetcher } from '@/actions/fetcher';
import { PaymentMethod } from '@/types';

const URL = `${process.env.NEXT_PUBLIC_FRONTEND_URL}/auth/paymentMethod`;

export const CreditCardSchema = yup.object().shape({
  number: yup
    .string()
    .required('Card number is required')
    .matches(/^[0-9]{13,16}$/, 'Card number must be 13 to 16 digits'),
  name: yup
    .string()
    .required('Name is required')
    .matches(/^[a-zA-Z ]*$/, 'Only alphabets are allowed for name')
    .min(2, 'Name must be at least 2 characters')
    .max(50, "Name can't be longer than 50 characters"),
  expiry: yup
    .string()
    .required('Expiry date is required')
    .matches(/^(0[1-9]|1[0-2])\/([0-9]{4}|[0-9]{2})$/, 'Must be a valid MM/YY format'),
  cvc: yup
    .string()
    .required('CVC is required')
    .matches(/^[0-9]{3,4}$/, 'Must be a valid CVC number'),
});

export interface PaymentMethodBody extends yup.InferType<typeof CreditCardSchema> {}

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
