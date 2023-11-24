import axios from "axios";
import qs from "query-string";

import { ProductItem } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/products`;

const getProductItems = async (id: string): Promise<ProductItem[]> => {
  const res = await axios.get(`${URL}/${id}/items/`);
  if (res?.data || res.data.data || res.data.data.length > 0) {
    return [];
  }
  return res.data.data[0];
};

export default getProductItems;
