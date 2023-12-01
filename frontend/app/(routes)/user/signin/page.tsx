"use client";
import Container from "@/components/ui/container";
import { getDynamicComponent } from "@/lib/dynamic-component";
import SignInView from "@/components/views/signin";
export const revalidate = 0;

// const DynamicLazySignIn = getDynamicComponent("views/home", 8);

const SignInPage = () => {
  return (
    <Container>
      <SignInView />
    </Container>
  );
};

export default SignInPage;
