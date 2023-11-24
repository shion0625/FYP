"use client";
import Skeleton, { SkeletonTheme } from "react-loading-skeleton";
import "react-loading-skeleton/dist/skeleton.css";

interface loadingSkeletonProps {
  baseColor?: string;
  highlightColor?: string;
  count?: number;
}
const LoadingSkeleton: React.FC<loadingSkeletonProps> = ({
  baseColor = "#202020",
  highlightColor = "#444",
  count = 4,
}) => {
  console.log("loading");
  return (
    <SkeletonTheme baseColor={baseColor} highlightColor={highlightColor}>
      <p>
        <Skeleton count={count} />
      </p>
    </SkeletonTheme>
  );
};

export default LoadingSkeleton;
