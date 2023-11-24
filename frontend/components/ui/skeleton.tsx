import Skeleton, { SkeletonTheme } from "react-loading-skeleton";

const ProductCardSkeleton = () => {
  return (
    <div className="bg-white group cursor-pointer rounded-xl border p-3 space-y-4">
      <div className="aspect-square rounded-xl bg-gray-100 relative">
        <Skeleton className="aspect-square object-cover rounded-md" />
        {/* Image part */}
      </div>
      {/* Description */}
      <div>
        <Skeleton className="text-lg" /> {/* name part */}
        <Skeleton className="text-sm " /> {/* description part */}
      </div>
      {/* Price */}
      <div className="flex itemx-center justify-between">
        <Skeleton /> {/* price part */}
      </div>
    </div>
  );
};

export const ProductCardListSkeleton = ({ count = 4 }) => {
  return (
    <div className="space-y-10 pb-10">
      <div className="flex flex-col gap-y-8 px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          {Array.from({ length: count }, (_, i) => (
            <ProductCardSkeleton key={i} />
          ))}
        </div>
      </div>
    </div>
  );
};
