'use client';
import SignUPView from '@/app/(routes)/user/signup/components/views';
import Container from '@/components/ui/container';
export const revalidate = 0;

const SignUpPage = () => (
  <Container>
    <SignUPView />
  </Container>
);

export default SignUpPage;
