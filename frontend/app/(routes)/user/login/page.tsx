"use client";
import Container from "@/components/ui/container";
import LoginInView from "@/components/views/login";
export const revalidate = 0;


const SignInPage = () => {
  return (
    <Container>
      <LoginInView />
    </Container>
  );
};

export default SignInPage;