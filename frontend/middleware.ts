import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { cookies } from "next/headers";
const COOKIE_SECURE = process.env.NEXT_PUBLIC_COOKIE_SECURE === "1";
const accessTokenExpiresInMinutes = 20;

export async function middleware(req: NextRequest) {
  const isBlockedRoot =
    req.nextUrl.pathname.startsWith("/cart") ||
    req.nextUrl.pathname.startsWith("/admin");
  const isProtectedAPI = req.nextUrl.pathname.startsWith("/api/auth");

  let res = NextResponse.next();
  if (isBlockedRoot) {
    res = await handlingAccessToken(res);
    if (res.ok) return res;
    return NextResponse.redirect(new URL("/user/login", req.url));
  }

  // if (isProtectedAPI) {
  //   res = await handlingAccessToken(res);
  //   return res;
  // }
  return res;
}

export const config = {
  matcher: ["/cart/:path*", "/admin/:path*", "/api/:path*"],
};

const handlingAccessToken = async (
  response: NextResponse<unknown>
): Promise<NextResponse<unknown>> => {
  // Add logic to check user login status
  const cookiesList = cookies();
  const hasAccessToken = cookiesList.has("access_token");
  const hasRefreshToken = cookiesList.has("refresh_token");

  if (hasAccessToken) return response;
  if (!hasRefreshToken) return respondWithSessionError();

  const refreshToken = cookiesList.get("refresh_token");
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_FRONTEND_URL}/renew-access-token`,
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
    return respondWithInternalServerError();
  }
  return response;
};

function respondWithInternalServerError() {
  return NextResponse.json(
    {
      error: "Internal Server Error",
      message: "renew access token is failed",
    },
    { status: 500 }
  );
}

function respondWithSessionError() {
  return NextResponse.json(
    {
      error: "Session has expired \n please login",
      message: "refreshToken is not exist",
    },
    { status: 401 }
  );
}
