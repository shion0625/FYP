import axios from "axios";
import qs from "query-string";

import { Product } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/products/`;

interface Query {
  categoryId?: string;
}

const getProducts = async (query: Query): Promise<Product[]> => {
  const url = qs.stringifyUrl({
    url: URL,
    query: {
      categoryId: query.categoryId,
    },
  });
  const res = await axios.get(url);
  return res.data.data[0];
};

export default getProducts;
