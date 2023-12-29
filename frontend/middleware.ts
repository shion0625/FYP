import { cookies } from 'next/headers';
import { NextResponse, type NextRequest } from 'next/server';

const COOKIE_SECURE = process.env.NEXT_PUBLIC_COOKIE_SECURE === '1';
const accessTokenExpiresInMinutes = 20;

export async function middleware(req: NextRequest) {
  const isBlockedRoot =
    req.nextUrl.pathname.startsWith('/cart') || req.nextUrl.pathname.startsWith('/admin');
  const isBlockedLoginRoot =
    req.nextUrl.pathname.startsWith('/user/login') ||
    req.nextUrl.pathname.startsWith('/user/signup');

  const isBlockedLogoutRoot =
    req.nextUrl.pathname.startsWith('/user/logout') || req.nextUrl.pathname.startsWith('/user');

  const isProtectedAPI = req.nextUrl.pathname.startsWith('/api/auth');

  const response = NextResponse.next();

  if (isBlockedRoot) {
    const res = await handlingAccessToken(response);
    if (res.ok) return response;
    return NextResponse.redirect(new URL('/user/login', req.url));
  }

  if (isBlockedLoginRoot) {
    const res = await handlingAccessToken(response);
    if (res.ok) {
      const redirectUrl = req.headers.get('referer') || '/';
      return NextResponse.redirect(new URL(redirectUrl, req.url));
    }
    return response;
  }

  if (isBlockedLogoutRoot) {
    const res = await handlingAccessToken(response);
    if (res.ok) {
      return response;
    }
    const redirectUrl = req.headers.get('referer') || '/';
    return NextResponse.redirect(new URL(redirectUrl, req.url));
  }

  if (isProtectedAPI) {
    const res = await handlingAccessToken(response);
    return res;
  }

  return response;
}

export const config = {
  matcher: ['/cart/:path*', '/api/:path*', '/user/login', '/user/signup', '/user/logout', '/user'],
};

const handlingAccessToken = async (
  response: NextResponse<unknown>
): Promise<NextResponse<unknown>> => {
  // Add logic to check user login status
  const cookiesList = cookies();
  const hasAccessToken = cookiesList.has('access_token');
  const hasRefreshToken = cookiesList.has('refresh_token');

  if (hasAccessToken) return response;
  if (!hasRefreshToken) return respondWithSessionError();

  const refreshToken = cookiesList.get('refresh_token');
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_FRONTEND_URL}/renew-access-token`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        refreshToken: refreshToken?.value,
      }),
    });
    const data = await res.json();
    if (data.data) {
      const expiryDate = new Date();
      expiryDate.setMinutes(expiryDate.getMinutes() + accessTokenExpiresInMinutes);
      response.cookies.set({
        name: 'access_token',
        value: data.data,
        httpOnly: true,
        path: '/',
        sameSite: 'strict',
        secure: COOKIE_SECURE,
        expires: expiryDate,
      });
    }
    return response;
  } catch (error) {
    return respondWithInternalServerError();
  }
};

function respondWithInternalServerError() {
  return NextResponse.json(
    {
      error: 'Internal Server Error',
      message: 'renew access token is failed',
    },
    { status: 500 }
  );
}

function respondWithSessionError() {
  return NextResponse.json(
    {
      error: 'Session has expired \n please login',
      message: 'refreshToken is not exist',
    },
    { status: 401 }
  );
}
