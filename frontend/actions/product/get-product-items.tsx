import axios from "axios";
import qs from "query-string";

import { ProductItem } from "@/types";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/products`;

const getProductItems = async (id: string): Promise<ProductItem[]> => {
  console.log("getProductItems" + id);
  const res = await axios.get(URL + "/" + id + "/items/");
  return res.data.data[0];
};

export default getProductItems;
