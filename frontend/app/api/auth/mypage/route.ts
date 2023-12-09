import { getAccessTokenCookie } from '@/utils/cookie';

const ProfileURL = `${process.env.NEXT_PUBLIC_API_URL}/account/`;
const AddressURL = `${process.env.NEXT_PUBLIC_API_URL}/account/address`;

export async function GET() {
  const accessToken = getAccessTokenCookie();
  // GoのAPIを呼び出す
  const responseProfile = await fetch(ProfileURL, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });
  const profileData = await responseProfile.json();

  const responseAddress = await fetch(AddressURL, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });
  const dataAddress = await responseAddress.json();

  // レスポンスを送信
  return Response.json({
    userProfile: profileData,
    userAddressList: dataAddress,
  });
}
