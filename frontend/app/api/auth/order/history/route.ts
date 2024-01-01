import { type NextRequest } from 'next/server';
import { getAccessTokenCookie } from '@/utils/cookie';

const URL = `${process.env.NEXT_PUBLIC_API_URL}/order/`;

export async function GET(req: NextRequest) {
  const accessToken = getAccessTokenCookie();

  const response = await fetch(URL, {
    method: req.method,
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });
  const data = await response.json();

  // レスポンスを送信
  return Response.json(data);
}
