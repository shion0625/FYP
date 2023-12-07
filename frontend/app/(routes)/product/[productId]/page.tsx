'use client';

import Container from '@/components/ui/container';
import { getDynamicComponent } from '@/lib/dynamic-component';
import { Loading } from './components/loading';

interface ProductPageProps {
  params: {
    productId: string;
  };
}

const DynamicLazyProductID = getDynamicComponent<ProductPageProps['params']>(
  'views/productID',
  <Loading count={16} />
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
