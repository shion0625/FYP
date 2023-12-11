import { getAccessTokenCookie } from '@/utils/cookie';
import type { NextRequest } from 'next/server';

const URL = `${process.env.NEXT_PUBLIC_API_URL}/account/payment-method`;


export async function GET() {
  const accessToken = getAccessTokenCookie();
  // GoのAPIを呼び出す
  const response = await fetch(URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });
  const data = await response.json();

  // レスポンスを送信
  return Response.json(data);
}

export async function POST(req: NextRequest) {
  const json = await req.json();
  const accessToken = getAccessTokenCookie();
  // GoのAPIを呼び出す
  const response = await fetch(URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify(json),
  });
  const data = await response.json();

  // レスポンスを送信
  return Response.json(data);
}
