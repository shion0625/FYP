import { ShoppingCart } from "lucide-react";

import { Product } from "@/types";
import Currency from "@/components/ui/currency";
import Button from "@/components/ui/button";
import ProductItemList from "@/components/product-item-list";
import getProductItems from "@/actions/product/get-product-items";

export const revalidate = 0;

interface InfoProps {
  data: Product;
}

const Info: React.FC<InfoProps> = async ({ data }) => {
  const productItems = await getProductItems(data.id);
  if (!productItems) {
    return null;
  }
  return (
    <div>
      <h1 className="text-3xl font-bold text-gray-900">{data.name}</h1>
      <div className="mt-3 flex items-end justify-between">
        <p className="text-2xl text-gray-900">
          <Currency value={data?.price} />
        </p>
      </div>
      <hr className="my-4" />
      <div className="flex flex-col gap-y-8 px-4 sm:px-6 lg:px-8">
        <ProductItemList title="Featured Products" items={productItems} />
      </div>
      <div className="mt-10 flex items-center gap-x-3">
        <Button className="flex items-center gap-x-2">
          Add To Cart
          <ShoppingCart />
        </Button>
      </div>
    </div>
  );
};

export default Info;
