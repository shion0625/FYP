'use client';
import { getDynamicComponent } from '@/lib/dynamic-component';
import Container from '@/components/ui/container';

export const revalidate = 0;

const DynamicLazyLogout = getDynamicComponent('user/logout', <></>);

const LogoutPage = () => (
  <Container>
    <DynamicLazyLogout />
  </Container>
);

export default LogoutPage;
