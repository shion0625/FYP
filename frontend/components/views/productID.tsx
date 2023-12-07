'use client';
import { useGetProduct } from '@/actions/product';
import Info from '@/components/info';
import ProductWrapper from '@/components/product/product-list-wrapper';
import NoResults from '@/components/ui/no-results';

interface ProductIDViewProps {
  productId: string;
}

const ProductIDView: React.FC<ProductIDViewProps> = ({ productId }) => {
  const { product } = useGetProduct(productId);

  return (
    <div className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mt-10 px-4 sm:mt-16 sm:px-0 lg:mt-0">
        {/* <ErrorBoundary> */}
        {product && product.data ? <Info data={product.data} /> : <NoResults />}
        {/* </ErrorBoundary> */}
      </div>
      <hr className="my-10" />
      <ProductWrapper product={product?.data} />
    </div>
  );
};

export default ProductIDView;
