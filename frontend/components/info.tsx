"use client";

import { ShoppingCart } from "lucide-react";

import { Product } from "@/types";
import Currency from "@/components/ui/currency";
import Button from "@/components/ui/button";
import ProductItemList from "@/components/product-item-list";
import { useGetProductItems } from "@/actions/product";
import NoResults from "@/components/ui/no-results";
import { toast } from "react-hot-toast";

export const revalidate = 0;

interface InfoProps {
  data: Product;
}

const Info: React.FC<InfoProps> = ({ data }) => {
  const { productItems, isError } = useGetProductItems(data.id);

  if (isError) {
    toast.error("Something went wrong.");
  }
  return (
    <div>
      <h1 className="text-3xl font-bold text-gray-900">{data.name}</h1>
      <div className="mt-3 flex items-end justify-between">
        <div className="text-2xl text-gray-900">
          <Currency value={data?.price} />
        </div>
      </div>
      <hr className="my-4" />
      <div className="flex flex-col gap-y-8 px-4 sm:px-6 lg:px-8">
        {productItems && productItems.data ? (
          <ProductItemList
            title="Featured Products"
            items={productItems.data}
          />
        ) : (
          <NoResults />
        )}
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
