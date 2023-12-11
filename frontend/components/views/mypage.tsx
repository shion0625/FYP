'use client';
import { Card } from 'flowbite-react';
import { UseGetMyPage } from '@/app/(routes)/user/myPage/hooks/get-mypage';
import CardIcon from '@/components/ui/credit-cards';
import NoResults from '@/components/ui/no-results';

const MyPageView = () => {
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
      {userAddressList && userAddressList.data ? (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 items-center bg-white shadow-lg rounded-lg p-10">
          <Card></Card>
          {userAddressList.data.map((address) => (
            <Card key={address.id}>
              <h2 className="text-2xl mb-4 font-semibold">{`Address ${address.id}`}</h2>
              <ul className="space-y-2 text-gray-700">
                <li>
                  <p>
                    <span className="font-bold">Name:</span> {address.name}
                  </p>
                </li>
                <li>
                  <p>
                    <span className="font-bold">Area:</span> {address.area}
                  </p>
                </li>
                <li>
                  <p>
                    <span className="font-bold">City:</span> {address.city}
                  </p>
                </li>
                <li>
                  <p>
                    <span className="font-bold">Country:</span> {address.countryName}
                  </p>
                </li>
                <li>
                  <p>
                    <span className="font-bold">House:</span> {address.house}
                  </p>
                </li>
                <li>
                  <p>
                    <span className="font-bold">Landmark:</span> {address.landMark}
                  </p>
                </li>
                <li>
                  <p>
                    <span className="font-bold">Phone Number:</span> {address.phoneNumber}
                  </p>
                </li>
                <li>
                  <p>
                    <span className="font-bold">Pincode:</span> {address.pincode}
                  </p>
                </li>
              </ul>
            </Card>
          ))}
        </div>
      ) : (
        <NoResults />
      )}
      {userPaymentMethod && userPaymentMethod.data ? (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 items-center bg-white shadow-lg rounded-lg p-10">
          {userPaymentMethod.data.map((paymentMethod) => {
            return (
              <Card key={paymentMethod.id}>
                <p>
                  <CardIcon cardCompany={paymentMethod.cardCompany} />
                  <span className="font-bold">Card ID:</span> {paymentMethod.id}
                </p>
                <p>
                  <span className="font-bold">Card Number:</span> **** **** ****{' '}
                  {paymentMethod.creditNumber}
                </p>
              </Card>
            );
          })}
        </div>
      ) : (
        <NoResults />
      )}
    </div>
  );
};

export default MyPageView;
