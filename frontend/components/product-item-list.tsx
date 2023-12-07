import { ProductItem } from '@/types';

import NoResults from '@/components/ui/no-results';
import ProductItemCard from '@/components/ui/product-item-card';

interface ProductListProps {
  title: string;
  items: ProductItem[];
}

const ProductItemList: React.FC<ProductListProps> = ({ title, items }) => {
  return (
    <div className="space-y-4">
      <h3 className="font-bold text-3xl">{title}</h3>
      {items.length === 0 && <NoResults />}
      <div className="grid lg:grid-cols-5 md:grid-cols-4 sm:grid-cols-3 max-sm:grid-cols-2 gap-8 font-mono text-white text-sm font-bold leading-6 bg-stripes-fuchsia rounded-lg text-center">
        {items.map((item) => (
          <ProductItemCard key={item.id} data={item} />
        ))}
      </div>
    </div>
  );
};

export default ProductItemList;
