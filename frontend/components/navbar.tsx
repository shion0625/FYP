import Link from "next/link";

import Container from "@/components/ui/container";
import NavbarActions from "@/components/navbar-actions";

export const revalidate = 0;

const Navbar = async () => {
  return (
    <header className="border-b">
      <Container>
        <div className="relative px-4 sm:px-6 lg:px-8 flex h-16 items-center">
          <Link href="/" className="ml-4 flex lg:ml-0 gap-x-2">
            <p className="font-bold text-xl">STORE</p>
          </Link>
          <NavbarActions />
        </div>
      </Container>
    </header>
  );
};

export default Navbar;
