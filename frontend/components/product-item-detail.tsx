"use client";

import { ProductItem } from "@/types";
import Currency from "@/components/ui/currency";
import Button from "@/components/ui/button";
import { useGetProductItems } from "@/actions/product";
import NoResults from "@/components/ui/no-results";
import { toast } from "react-hot-toast";
import React, { useState, useMemo } from "react";
import Variation from "@/components/variation";

export const revalidate = 0;

interface ProductItemDetailProps {
  data: ProductItem;
}

const ProductItemDetail: React.FC<ProductItemDetailProps> = ({ data }) => {
  // nameの配列（重複なし）
  const names = useMemo(
    () => Array.from(new Set(data.variationValues.map((item) => item.name))),
    [data.variationValues]
  );

  // valueのmap配列(keyがname)
  const valuesMap = useMemo(
    () =>
      data.variationValues.reduce((acc: { [key: string]: string[] }, curr) => {
        if (!acc[curr.name]) {
          acc[curr.name] = [];
        }
        acc[curr.name].push(curr.value);
        return acc;
      }, {}),
    [data.variationValues]
  );

  // 選択された値を追跡するためのstate
  const [selectedValues, setSelectedValues] = useState<{
    [key: string]: string | null;
  }>({});

  return (
    <div>
      <div className="mt-3 flex items-end justify-between">
        <h2 className="text-3xl font-bold text-gray-900">{data.name}</h2>
        <div className="text-2xl text-gray-900">
          <Currency
            value={data?.price}
            discountPrice={data?.discountPrice}
          />
        </div>
      </div>
      <hr className="my-4" />
      <div>
        <h3 className="text-3xl font-bold text-gray-900">variation</h3>
        {names.map((name, index) => (
          <Variation
            key={index}
            name={name}
            values={valuesMap[name]}
            selectedValue={selectedValues[name]}
            onSelect={(value) =>
              setSelectedValues((prev) => ({ ...prev, [name]: value }))
            }
          />
        ))}
      </div>
    </div>
  );
};

export default ProductItemDetail;
