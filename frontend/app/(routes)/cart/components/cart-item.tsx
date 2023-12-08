import React from 'react';
import { X } from 'lucide-react';
import Gallery from '@/components/gallery';
import Currency from '@/components/ui/currency';
import IconButton from '@/components/ui/icon-button';
import useCart from '@/hooks/use-cart';
import { ProductItem } from '@/types';

interface CartItemProps {
  data: ProductItem;
}

const CartItem: React.FC<CartItemProps> = ({ data }) => {
  const cart = useCart();

  const onRemove = () => {
    cart.removeItem(data.id);
  };

  return (
    <li className="flex py-6 border-b">
      <div className="relative h-24 w-24 rounded-md overflow-hidden sm:h-48 sm:w-48">
        <Gallery id={data.sku} urls={data.images} />
      </div>
      <div className="relative ml-4 flex flex-1 flex-col justify-between sm:ml-6">
        <div className="absolute z-10 right-0 top-0">
          <IconButton onClick={onRemove} icon={<X size={15} />} />
        </div>
        <div className="relative pr-9 sm:grid sm:grid-cols-2 sm:gap-x-6 sm:pr-0">
          <div className="flex justify-between">
            <p className=" text-lg font-semibold text-black">{data.name}</p>
          </div>

          <div className="mt-1 flex text-sm">
            {data.variationValues.map((variation) => (
              <p key={data.name + variation.variationId} className="text-gray-500">
                {variation.name}: {variation.value}
              </p>
            ))}
          </div>
          <Currency value={data.price} />
        </div>
      </div>
    </li>
  );
};

export default CartItem;
