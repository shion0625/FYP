import { useGetProduct, useGetProducts } from "@/actions/product";
import LoadingSkeleton from "@/components/ui/loading-skeleton";
import ProductList from "@/components/product-list";
import dynamic from "next/dynamic";
import NoResults from "@/components/ui/no-results";

const DynamicLazyInfo = dynamic(() => import("@/components/info"), {
  ssr: false,
  loading: () => <LoadingSkeleton />,
});

interface ProductIDViewProps {
  productId: string;
}

const ProductIDView: React.FC<ProductIDViewProps> = ({ productId }) => {
  const { product } = useGetProduct(productId);

  const { products } = useGetProducts({
    categoryId: product?.data?.categoryId,
  });
  return (
    <div className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mt-10 px-4 sm:mt-16 sm:px-0 lg:mt-0">
        {/* <ErrorBoundary> */}
        {product && product.data ? (
          <DynamicLazyInfo data={product.data} />
        ) : (
          <NoResults />
        )}
        {/* </ErrorBoundary> */}
      </div>
      <hr className="my-10" />
      {products && products.data ? (
        <ProductList title="Related Items" items={products?.data} />
      ) : (
        <NoResults />
      )}
    </div>
  );
};

export default ProductIDView;
