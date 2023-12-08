'use client';

import { getDynamicComponent } from '@/lib/dynamic-component';
import Container from '@/components/ui/container';
import { Loading } from './components/loading';

export const revalidate = 0;

interface CategoryIdPageProps {
  params: {
    categoryId: string;
  };
}
const DynamicLazyCategoryID = getDynamicComponent<CategoryIdPageProps['params']>(
  'views/categoryID',
  <Loading count={16} />
);

const CategoryIdPage = ({ params }: CategoryIdPageProps) => (
  <Container>
    <DynamicLazyCategoryID categoryId={params.categoryId} />
  </Container>
);

export default CategoryIdPage;
