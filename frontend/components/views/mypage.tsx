'use client';
import { Card } from 'flowbite-react';
import { useGetAllAddresses } from '@/actions/user/user-address';
import { useGetProfile } from '@/actions/user/user-profile';
import NoResults from '@/components/ui/no-results';
import useUserId from '@/hooks/use-user-profile';

const MyPageView = () => {
  const userId = useUserId();
  const { userProfile } = useGetProfile({ userId: userId.userId });
  const { userAddressList } = useGetAllAddresses({ userId: userId.userId });
  console.log(userAddressList);

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
          {userAddressList.data.map((address, index) => (
            <Card key={index}>
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
    </div>
  );
};

export default MyPageView;
