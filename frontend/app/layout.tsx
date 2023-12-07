import type { Metadata } from 'next';

import { Urbanist } from 'next/font/google';

import { getCategories } from '@/actions/category';
import ModalProvider from '@/providers/modal-provider';
import ToastProvider from '@/providers/toast-provider';

import Footer from '@/components/layout/footer';
import Navbar from '@/components/layout/navbar';
import Sidebar from '@/components/layout/sidebar';

import './globals.css';

const font = Urbanist({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Store',
  description: 'Store',
};

export default async function RootLayout({ children }: { children: React.ReactNode }) {
  const categories = await getCategories();

  return (
    <html lang="en">
      <body className={`${font.className} vsc-initialized`}>
        <ToastProvider />
        <ModalProvider />
        <Navbar />
        <Sidebar data={categories?.data} />
        {children}
        <Footer />
      </body>
    </html>
  );
}
