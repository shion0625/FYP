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

export const ProductItemCardSkelton = () => {
  return (
    <div className="bg-white group cursor-pointer rounded-xl border p-3 space-y-4">
      <div className="aspect-square rounded-xl bg-gray-100 relative">
        <div className="aspect-square rounded-xl bg-gray-100 relative mb-4">
          <Skeleton className="aspect-square object-cover rounded-md" />{" "}
        </div>
      </div>
      {/* Name */}
      <div>
        <h2 className="text-lg text-stone-800">
          <Skeleton className="text-lg" width={"50%"} />
        </h2>
      </div>
      {/* Price */}
      <Skeleton className="text-2xl" width={"40%"} />
      {/* Description */}
      <div>
        <Skeleton className="text-sm" width={"50%"} />
      </div>
    </div>
  );
};

export const ProductItemCardListSkeleton = ({ count = 4 }) => {
  return (
    <div className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mt-10 px-4 sm:mt-16 sm:px-0 lg:mt-0">
        <div>
          <Skeleton className="text-3xl" width={"60%"} />
          <div className="mt-3">
            <div className="text-2xl text-gray-900">
              <Skeleton className="text-2xl" width={"40%"} />
            </div>
          </div>
          <hr className="my-4" />
          <div className="flex flex-col gap-y-8 px-4 sm:px-6 lg:px-8">
            <div className="space-y-4">
              <h3 className="font-bold text-3xl">
                <Skeleton className="text-3xl" width={"60%"} />
              </h3>
              <div className="grid lg:grid-cols-5 md:grid-cols-4 sm:grid-cols-3 max-sm:grid-cols-2 gap-8 font-mono text-white text-sm font-bold leading-6 bg-stripes-fuchsia rounded-lg text-center">
                {Array.from({ length: count }, (_, i) => (
                  <ProductItemCardSkelton key={i} />
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>
      <hr className="my-10" />
      <div className="space-y-4">
        <h3 className="font-bold text-3xl">
          <Skeleton className="text-3xl" width={"60%"} />
        </h3>
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          {Array.from({ length: count }, (_, i) => (
            <ProductCardSkeleton key={i} />
          ))}
        </div>
      </div>
    </div>
  );
};
