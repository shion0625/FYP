// import useSWR from "swr";
// import { axiosFetcher } from "@/actions/fecher";
import axios from "axios";
import { Response, Category } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/categories/`;

interface UseGetProductItemsReturn {
  categories?: Response<Category[]>;
  isError: any;
}

export const getCategories = async (): Promise<
  UseGetProductItemsReturn["categories"]
> => {
  const res = await axios.get(URL);
  return res.data;
};

// export const useGetCategories = (): UseGetProductItemsReturn => {
//   const { data, error } = useSWR<UseGetProductItemsReturn["categories"]>(
//     URL,
//     axiosFetcher,
//     {
//       suspense: true,
//     }
//   );

//   return {
//     categories: data,
//     isError: error,
//   };
// };
