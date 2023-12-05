import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { cookies } from "next/headers";
const COOKIE_SECURE =
  process.env.NEXT_PUBLIC_COOKIE_SECURE == "1" ? true : false;
const accessTokenExpiresInMinutes = 20;

export async function middleware(req: NextRequest) {
  let res = NextResponse.next();
  if (
    req.nextUrl.pathname.startsWith("/cart") ||
    req.nextUrl.pathname.startsWith("/admin")
  ) {
    res = await handlingAccessToken(res);
    if (res.ok) {
      return res;
    }
    // ログインしていないユーザーをログインページにリダイレクト
    return NextResponse.redirect(new URL("/user/login", req.url));
  }

  if (req.nextUrl.pathname.startsWith("/api/cart")) {
    res = await handlingAccessToken(res);
    return res;
  }
  return res;
}

export const config = {
  matcher: ["/cart/:path*", "/admin/:path*", "/api/:path*"],
};

const handlingAccessToken = async (
  response: NextResponse<unknown>
): Promise<NextResponse<unknown>> => {
  // ユーザーのログイン状態をチェックするロジックを追加
  const cookiesList = cookies();
  const hasAccessToken = cookiesList.has("access_token");
  const hasRefreshToken = cookiesList.has("refresh_token");
  if (hasAccessToken) {
    return response;
  }
  if (!hasRefreshToken) {
    return NextResponse.json(
      {
        error: "Session has expired \n please login",
        message: "refreshToken is not exist",
      },
      { status: 401 }
    );
  }

  const refreshToken = cookiesList.get("refresh_token");
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_FRONTEND_URL}/user/renew-access-token`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          refreshToken: refreshToken?.value,
        }),
      }
    );
    const data = await res.json();
    const expiryDate = new Date();
    expiryDate.setMinutes(
      expiryDate.getMinutes() + accessTokenExpiresInMinutes
    );
    response.cookies.set({
      name: "access_token",
      value: data.data,
      httpOnly: true,
      path: "/",
      sameSite: "strict",
      secure: COOKIE_SECURE,
      expires: expiryDate,
    });
  } catch (error) {
    return NextResponse.json(
      {
        error: "Internal Server Error",
        message: "renew access token is failed",
      },
      { status: 500 }
    );
  }
  return response;
};
