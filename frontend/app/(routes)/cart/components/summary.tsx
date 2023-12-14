'use client';
import React, { useEffect, useState, useRef } from 'react';
import { toast } from 'react-hot-toast';
import { Button, Card } from 'flowbite-react';
import { useRouter } from 'next/navigation';
import { UsePurchase } from '@/actions/order/purchase';
import { UsePaymentMethod } from '@/actions/user/payment-method';
import { UseUserAddress } from '@/actions/user/user-address';
import CreditCardsForm from '@/components/credit-cards-form';
import BackdropModal from '@/components/ui/backdrop-modal';
import CardIcon from '@/components/ui/credit-cards';
import Currency from '@/components/ui/currency';
import useCart from '@/hooks/use-cart';
import { Address, PaymentMethod } from '@/types';

interface DataState {
  paymentMethod: PaymentMethod[];
  userAddress: Address[];
}
const Summary = () => {
  const router = useRouter();
  const items = useCart((state) => state.items);
  const removeAll = useCart((state) => state.removeAll);
  const { purchaseOrder } = UsePurchase();
  const { getUserAddresses } = UseUserAddress();
  const { getPaymentMethod } = UsePaymentMethod();
  const [responseData, setData] = useState<DataState>({
    paymentMethod: [],
    userAddress: [],
  });

  useEffect(() => {
    const fetchData = async () => {
      const paymentMethod = await getPaymentMethod();
      const userAddress = await getUserAddresses();
      setData({ paymentMethod, userAddress });
    };
    fetchData();
  }, []);

  const totalPrice = items.reduce((total, item) => total + Number(item.price), 0);

  const clickedAddressId = useRef(0);

  const handleCardClick = (id: number) => {
    clickedAddressId.current = id;
  };

  const onCheckout = async () => {
    const convertProductItemInfo = items.map((item) => ({
      productItemId: item.id,
      variationValues: item.variationValues,
      count: item.count || 0, // Countがundefinedの場合は0とする
    }));
    try {
      const response = await purchaseOrder({
        addressId: clickedAddressId.current,
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
      <div className="grid grid-cols-1 items-center bg-white shadow-lg rounded-lg p-10 my-4">
        <Button
          color="white"
          className="p-5 border-dashed border-2 border-gray-500 cursor-pointer hover:bg-gray-200 transition-colors duration-200 ease-in-out"
          onClick={() => router.push('/user/address/add')}
        >
          add address
        </Button>
        {responseData.userAddress &&
          responseData.userAddress.length > 0 &&
          responseData.userAddress.map((address, index) => (
            <Card className="mt-4" key={index} onClick={() => handleCardClick(address.id)}>
              <h2 className="text-2xl mb-4 font-semibold">{address.name}</h2>
              <p className="space-y-2 text-gray-700">
                {address.house}
                <br />
                {address.city}, {address.area}, {address.pincode}
                <br />
                {address.countryName}
                <br />
                TEL: {address.phoneNumber}
                <br />
                landMark: {address.landMark}
              </p>
            </Card>
          ))}
      </div>
      <div className="grid grid-cols-1 items-center bg-white shadow-lg rounded-lg p-10 mb-4">
        <BackdropModal
          buttonClassName="p-5 border-dashed border-2 border-gray-500 cursor-pointer hover:bg-gray-200 transition-colors duration-200 ease-in-out"
          buttonText="add payment method"
          headerText="Add Payment Method"
        >
          <CreditCardsForm />
        </BackdropModal>
        {responseData.paymentMethod &&
          responseData.paymentMethod.length > 0 &&
          responseData.paymentMethod.map((paymentMethod) => (
            <Card className="mt-4" key={paymentMethod.id}>
              <p>
                <CardIcon cardCompany={paymentMethod.cardCompany} />
                <span className="font-bold">Card ID:</span> {paymentMethod.id}
              </p>
              <p>
                <span className="font-bold">Card Number:</span> **** **** ****{' '}
                {paymentMethod.number}
              </p>
            </Card>
          ))}
      </div>
      <Button onClick={onCheckout} disabled={items.length === 0} className="w-full mt-6">
        Checkout
      </Button>
    </div>
  );
};

export default Summary;
