'use client';

import { getDynamicComponent } from '@/lib/dynamic-component';
import Container from '@/components/ui/container';
import { Loading } from './components/loading';

interface ProductPageProps {
  params: {
    productId: string;
  };
}

const DynamicLazyProductID = getDynamicComponent<ProductPageProps['params']>(
  'product/[productId]',
  <Loading count={16} />
);

const ProductPage: React.FC<ProductPageProps> = ({ params }) => (
  <div className="bg-white">
    <Container>
      <DynamicLazyProductID productId={params.productId} />
    </Container>
  </div>
);

export default ProductPage;
