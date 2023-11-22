import axios from "axios";
import qs from "query-string";

import { Product } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/products/`;

interface Query {
  categoryId?: string;
  brandId?: string;
}

const getProducts = async (query: Query): Promise<Product[]> => {
  const url = qs.stringifyUrl({
    url: URL,
    query: {
      category_id: query.categoryId,
      brand_id: query.brandId,
    },
  });

  const res = await axios.get(url);
  return res.data.data[0];
};

export default getProducts;
