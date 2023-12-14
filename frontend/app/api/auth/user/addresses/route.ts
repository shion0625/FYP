import { getAccessTokenCookie } from '@/utils/cookie';

const URL = `${process.env.NEXT_PUBLIC_API_URL}/account/addresses`;

export async function GET() {
  const accessToken = getAccessTokenCookie();
  // GoのAPIを呼び出す
  const response = await fetch(URL, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });
  const data = await response.json();

  // レスポンスを送信
  return Response.json(data);
}
