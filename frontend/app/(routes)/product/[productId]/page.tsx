"use client";

import Container from "@/components/ui/container";
import { getDynamicComponent } from "@/lib/dynamic-component";
import { ProductCardListSkeleton } from "@/components/ui/skeleton";

interface ProductPageProps {
  params: {
    productId: string;
  };
}

const DynamicLazyProductID = getDynamicComponent<ProductPageProps["params"]>(
  "views/productID",
  8
);

const ProductPage: React.FC<ProductPageProps> = ({ params }) => {
  return (
    <div className="bg-white">
      <Container>
        <DynamicLazyProductID productId={params.productId} />
      </Container>
    </div>
  );
};

export default ProductPage;
