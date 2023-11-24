import getProduct from "@/actions/product/get-product";
import getProducts from "@/actions/product/get-products";
import LoadingSkeleton from "@/components/ui/loading-skeleton";
import ProductList from "@/components/product-list";
import Container from "@/components/ui/container";
import ErrorBoundary from "@/lib/error-boundary";
import dynamic from "next/dynamic";

const DynamicLazyInfo = dynamic(() => import("@/components/info"), {
  ssr: false,
  loading: () => <LoadingSkeleton />,
});

interface ProductPageProps {
  params: {
    productId: string;
  };
}

const ProductPage: React.FC<ProductPageProps> = async ({ params }) => {
  const product = await getProduct(params.productId);
  const suggestedProducts = await getProducts({
    categoryId: product?.categoryId,
  });

  if (!product) {
    return null;
  }

  return (
    <div className="bg-white">
      <Container>
        <div className="px-4 py-10 sm:px-6 lg:px-8">
          <div className="mt-10 px-4 sm:mt-16 sm:px-0 lg:mt-0">
            {/* <ErrorBoundary> */}
            <DynamicLazyInfo data={product} />
            {/* </ErrorBoundary> */}
          </div>
          <hr className="my-10" />
          <ProductList title="Related Items" items={suggestedProducts} />
        </div>
      </Container>
    </div>
  );
};

export default ProductPage;
