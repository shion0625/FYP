'use client';

import { toast } from 'react-hot-toast';
import { useGetProductItems } from '@/actions/product';
import ProductItemList from '@/components/product/product-item-list';
import NoResults from '@/components/ui/no-results';
import { Product } from '@/types';

export const revalidate = 0;

interface InfoProps {
  data: Product;
}

const Info: React.FC<InfoProps> = ({ data }) => {
  const { productItems, isError } = useGetProductItems(data.id);

  if (isError) {
    toast.error('Something went wrong.');
  }
  return (
    <div>
      <h1 className="text-3xl font-bold text-gray-900">{data.name}</h1>
      <div className="mt-3 flex items-end justify-between"></div>
      <hr className="my-4" />
      <div className="flex flex-col gap-y-8 px-4 sm:px-6 lg:px-8">
        {productItems && productItems.data ? (
          <ProductItemList title="Featured Products" items={productItems.data} />
        ) : (
          <NoResults />
        )}
      </div>
    </div>
  );
};

export default Info;
