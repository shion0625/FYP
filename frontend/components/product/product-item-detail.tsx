'use client';

import React from 'react';
import Currency from '@/components/ui/currency';
import Variation from '@/components/variation';
import { ProductItem, ProductVariationValue } from '@/types';

export const revalidate = 0;

interface ProductItemDetailProps {
  data: ProductItem;
  names: string[];
  variationsMap: {
    [key: string]: ProductVariationValue[];
  };
  selectedValues: {
    [key: string]: ProductVariationValue | null;
  };
  setSelectedValues: React.Dispatch<
    React.SetStateAction<{
      [key: string]: ProductVariationValue | null;
    }>
  >;
}

const ProductItemDetail: React.FC<ProductItemDetailProps> = ({
  data,
  names,
  variationsMap,
  selectedValues,
  setSelectedValues,
}) => (
  <div>
    <div className="mt-3 flex items-end justify-between">
      <h2 className="text-3xl font-bold text-gray-900">{data.itemName}</h2>
      <div className="text-2xl text-gray-900">
        <Currency value={data?.price} discountPrice={data?.discountPrice} />
      </div>
    </div>
    <hr className="my-4" />
    <div>
      <h3 className="text-3xl font-bold text-gray-900">variation</h3>
      {names.map((name, index) => (
        <Variation
          key={index}
          name={name}
          productVariationValues={variationsMap[name]}
          selectedValue={selectedValues[name]}
          onSelect={(value) => setSelectedValues((prev) => ({ ...prev, [name]: value }))}
        />
      ))}
    </div>
  </div>
);

export default ProductItemDetail;
