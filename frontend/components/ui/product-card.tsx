'use client';

import NextImage from 'next/image';
import { useRouter } from 'next/navigation';

import { Product } from '@/types';
import Currency from '@/components/ui/currency';

interface ProductCard {
  data: Product;
}

const ProductCard: React.FC<ProductCard> = ({ data }) => {
  const router = useRouter();

  const handleClick = () => {
    router.push(`/product/${data?.id}`);
  };

  return (
    <div
      onClick={handleClick}
      className="bg-white group cursor-pointer rounded-xl border p-3 space-y-4"
    >
      {/* Images and Actions */}
      <div className="aspect-square rounded-xl bg-gray-100 relative">
        <NextImage
          src={data?.image}
          fill
          sizes="(max-width: 600px) 100vw, 600px"
          alt="NextImage"
          className="aspect-square object-cover rounded-md"
        />
      </div>
      {/* Description */}
      <div>
        <p className="font-semibold text-lg">{data.name}</p>
        <p className="text-sm text-gray-500">{data.description}</p>
      </div>
      {/* Price */}
      <div className="flex itemx-center justify-between">
        <Currency value={data?.price} />
      </div>
    </div>
  );
};

export default ProductCard;
