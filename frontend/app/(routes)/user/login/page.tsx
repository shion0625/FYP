'use client';
import Container from '@/components/ui/container';
import LoginInView from '@/app/(routes)/user/login/components/views';
export const revalidate = 0;

const LoginPage = () => (
  <Container>
    <LoginInView />
  </Container>
);

export default LoginPage;
