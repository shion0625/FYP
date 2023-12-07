'use client';

import Container from '@/components/ui/container';
import { getDynamicComponent } from '@/lib/dynamic-component';

interface ProductPageProps {
  params: {
    productId: string;
  };
}

const DynamicLazyProductID = getDynamicComponent<ProductPageProps['params']>(
  'views/productID',
  8,
  'productCardItem'
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
