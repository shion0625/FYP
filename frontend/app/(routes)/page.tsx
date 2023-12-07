'use client';
import Container from '@/components/ui/container';
import { getDynamicComponent } from '@/lib/dynamic-component';

export const revalidate = 0;

const DynamicLazyHome = getDynamicComponent('views/home', 8);

const HomePage = () => {
  return (
    <Container>
      <DynamicLazyHome />
    </Container>
  );
};

export default HomePage;
