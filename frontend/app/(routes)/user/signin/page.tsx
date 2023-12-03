"use client";
import Container from "@/components/ui/container";
import SignInView from "@/components/views/signin";
export const revalidate = 0;

const SignInPage = () => {
  return (
    <Container>
      <SignInView />
    </Container>
  );
};

export default SignInPage;