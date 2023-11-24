"use client";
import { SkeletonTheme } from "react-loading-skeleton";
import { ProductCardListSkeleton } from "@/components/ui/skeleton";
import "react-loading-skeleton/dist/skeleton.css";

interface loadingSkeletonProps {
  baseColor?: string;
  highlightColor?: string;
  count?: number;
  type?: "productCard" | "productCardItem";
}
const LoadingSkeleton: React.FC<loadingSkeletonProps> = ({
  baseColor = "#ebebeb",
  highlightColor = "#f5f5f5",
  count = 4,
  type = "productCard",
}) => {
  return (
    <SkeletonTheme baseColor={baseColor} highlightColor={highlightColor}>
      {type == "productCard" && <ProductCardListSkeleton count={count} />}
    </SkeletonTheme>
  );
};

export default LoadingSkeleton;
