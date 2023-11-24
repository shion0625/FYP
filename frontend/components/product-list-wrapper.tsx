"use client";
import { useGetProduct, useGetProducts } from "@/actions/product";
import LoadingSkeleton from "@/components/ui/loading-skeleton";
import ProductList from "@/components/product-list";
import dynamic from "next/dynamic";
import NoResults from "@/components/ui/no-results";
import { Product } from "@/types";

interface ProductIDViewProps {
  product?: Product | null;
}

const ProductWrapper: React.FC<ProductIDViewProps> = ({ product }) => {
  const { products } = useGetProducts({
    categoryId: product?.categoryId,
  });

  return (
    <>
      {products && products.data ? (
        <ProductList title="Related Items" items={products?.data} />
      ) : (
        <NoResults />
      )}
    </>
  );
};

export default ProductWrapper;
