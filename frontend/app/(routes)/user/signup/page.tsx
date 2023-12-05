"use client";
import Container from "@/components/ui/container";
import SignUPView from "@/components/views/signup";
export const revalidate = 0;

const SignUpPage = () => {
  return (
    <Container>
      <SignUPView />
    </Container>
  );
};

export default SignUpPage;
