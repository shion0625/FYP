import type { NextRequest } from 'next/server';
import { setAccessTokenCookie, setRefreshTokenCookie } from '@/utils/cookie';
const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/login/`;

export async function POST(req: NextRequest) {
  const json = await req.json();

  // Call the Go API
  const response = await fetch(URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(json),
  });
  const data = await response.json();

  // Get the access token
  const { accessToken, refreshToken } = data.data;

  // Set the access token in the cookie
  setAccessTokenCookie(accessToken);
  setRefreshTokenCookie(refreshToken);
  // Send the response
  return Response.json(data);
}
