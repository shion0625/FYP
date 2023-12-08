'use client';
import { getDynamicComponent } from '@/lib/dynamic-component';
import Container from '@/components/ui/container';
import { Loading } from './components/loading';

export const revalidate = 0;

const DynamicLazyHome = getDynamicComponent('views/home', <Loading count={16} />);

const HomePage = () => (
  <Container>
    <DynamicLazyHome />
  </Container>
);

export default HomePage;
