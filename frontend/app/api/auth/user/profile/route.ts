import { NextResponse, type NextRequest } from 'next/server';
import { getAccessTokenCookie } from '@/utils/cookie';

const URL = `${process.env.NEXT_PUBLIC_API_URL}/account/`;

export async function GET(req: NextRequest) {
  const accessToken = getAccessTokenCookie();
  // GoのAPIを呼び出す
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

export async function PUT(req: NextRequest) {
  const json = await req.json();
  const accessToken = getAccessTokenCookie();
  // GoのAPIを呼び出す
  const response = await fetch(URL, {
    method: req.method,
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify(json),
  });
  if (!response.ok) {
    throw new Error(`Server responded with status ${response.status}`);
  }

  // レスポンスを送信
  return NextResponse.json(response);
}
