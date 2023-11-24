"use client";
import Link from "next/link";
import useSidebar from "@/hooks/use-sidebar";
import Container from "@/components/ui/container";
import NavbarActions from "@/components/navbar-actions";
import HamburgerMenu from "@/components/ui/hamburger-button";
export const revalidate = 0;

const Navbar = () => {
  const sidebar = useSidebar();
  return (
    <header className="border-b">
      <Container>
        <div className="relative px-4 sm:px-6 lg:px-8 flex h-16 items-center">
          <HamburgerMenu
            isOpen={sidebar.isOpen}
            onOpen={sidebar.onOpen}
            onClose={sidebar.onClose}
          />
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
