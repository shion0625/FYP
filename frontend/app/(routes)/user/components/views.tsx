'use client';
import { Button, Card } from 'flowbite-react';
import { useRouter } from 'next/navigation';
import { UseGetMyPage } from '@/app/(routes)/user/hooks/get-mypage';
import CreditCardsForm from '@/components/credit-cards-form';
import BackdropModal from '@/components/ui/backdrop-modal';
import CardIcon from '@/components/ui/credit-cards';
import NoResults from '@/components/ui/no-results';

const MyUserView = () => {
  const router = useRouter();
  const { userProfile, userAddressList, userPaymentMethod } = UseGetMyPage();
  return (
    <div className="space-y-10 pb-10">
      {userProfile && userProfile.data ? (
        <div className="flex flex-col items-center bg-white shadow-lg rounded-lg p-10">
          <h1 className="text-4xl mb-4 font-semibold">{`${userProfile.data.firstName} ${userProfile.data.lastName}`}</h1>
          <ul className="space-y-2 text-gray-700">
            <li>
              <p>
                <span className="font-bold">Username:</span> {userProfile.data.userName}
              </p>
            </li>
            <li>
              <p>
                <span className="font-bold">Age:</span> {userProfile.data.age}
              </p>
            </li>
            <li>
              <p>
                <span className="font-bold">Email:</span> {userProfile.data.email}
              </p>
            </li>
          </ul>
        </div>
      ) : (
        <NoResults />
      )}
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 items-center bg-white shadow-lg rounded-lg p-10">
        <Button
          color="white"
          className="h-full border-dashed border-2 border-gray-500 cursor-pointer hover:bg-gray-200 transition-colors duration-200 ease-in-out"
          onClick={() => router.push('/user/address/add')}
        >
          add address
        </Button>
        {userAddressList?.data &&
          userAddressList.data.length > 0 &&
          userAddressList.data.map((address) => (
            <Card key={address.id}>
              <h2 className="text-2xl mb-4 font-semibold">{address.name}</h2>
              <p className="space-y-2 text-gray-700">
                {address.house}
                <br />
                {address.city},{address.area}, {address.pincode}
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
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 items-center bg-white shadow-lg rounded-lg p-10">
        <BackdropModal
          buttonClassName="h-full border-dashed border-2 border-gray-500 cursor-pointer hover:bg-gray-200 transition-colors duration-200 ease-in-out"
          buttonText="add payment method"
          headerText="Add Payment Method"
        >
          <CreditCardsForm />
        </BackdropModal>
        {userPaymentMethod?.data &&
          userPaymentMethod?.data.length > 0 &&
          userPaymentMethod.data.map((paymentMethod) => {
            return (
              <Card key={paymentMethod.id}>
                <p>
                  <CardIcon cardCompany={paymentMethod.cardCompany} />
                  <span className="font-bold">Card ID:</span> {paymentMethod.id}
                </p>
                <p>
                  <span className="font-bold">Card Number:</span> **** **** ****{' '}
                  {paymentMethod.number}
                </p>
              </Card>
            );
          })}
      </div>
    </div>
  );
};

export default MyUserView;
