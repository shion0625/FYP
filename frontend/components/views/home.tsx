'use client';

import { useGetProducts } from '@/actions/product';
import ProductList from '@/components/product/product-list';
import NoResults from '@/components/ui/no-results';

const HomeView = () => {
  const { products } = useGetProducts({});
  return (
    <div className="space-y-10 pb-10">
      {/* <Billboard data={billboard} /> */}
      <div className="flex flex-col gap-y-8 px-4 sm:px-6 lg:px-8">
        {products && products.data ? (
          <ProductList title="Featured Products" items={products.data} />
        ) : (
          <NoResults />
        )}
      </div>
    </div>
  );
};

export default HomeView;
