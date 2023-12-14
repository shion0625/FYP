'use client';
import LoginInView from '@/app/(routes)/user/login/components/views';
import Container from '@/components/ui/container';
export const revalidate = 0;

const LoginPage = () => (
  <Container>
    <LoginInView />
  </Container>
);

export default LoginPage;
