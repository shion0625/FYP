'use client';

import { toast } from 'react-hot-toast';
import { Button } from 'flowbite-react';
import { UsePurchase } from '@/actions/cart/purchase';
import Currency from '@/components/ui/currency';
import useCart from '@/hooks/use-cart';

const Summary = () => {
  const items = useCart((state) => state.items);
  const removeAll = useCart((state) => state.removeAll);
  const { purchaseOrder } = UsePurchase();

  const totalPrice = items.reduce((total, item) => total + Number(item.price), 0);

  const onCheckout = async () => {
    const convertProductItemInfo = items.map((item) => ({
      productItemId: item.id,
      variationValues: item.variationValues,
      count: item.count || 0, // Countがundefinedの場合は0とする
    }));
    try {
      const response = await purchaseOrder({
        userId: '016e5917-938b-4168-bc37-3da46f9f8343',
        addressId: 11,
        productItemInfo: convertProductItemInfo,
        totalFee: totalPrice,
        paymentMethodID: 11,
      });
      toast.success(response.message);
      removeAll();
    } catch (error: unknown) {
      toast.error('failed to purchase');
    }
  };

  return (
    <div className="mt-16 rounded-lg bg-gray-50 px-4 py-6 sm:p-6 lg:col-span-5 lg:mt-0 lg:p-8">
      <h2 className="text-lg font-medium text-gray-900">Order summary</h2>
      <div className="mt-6 space-y-4">
        <div className="flex items-center justify-between border-t border-gray-200 pt-4">
          <div className="text-base font-medium text-gray-900">Order total</div>
          <Currency value={totalPrice} />
        </div>
      </div>
      <Button onClick={onCheckout} disabled={items.length === 0} className="w-full mt-6">
        Checkout
      </Button>
    </div>
  );
};

export default Summary;
