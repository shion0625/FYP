import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { cookies } from "next/headers";

export function middleware(req: NextRequest) {
  const res = NextResponse.next();

  if (
    req.nextUrl.pathname.startsWith("/cart") ||
    req.nextUrl.pathname.startsWith("/admin")
  ) {
    if (handlingAccessToken()) {
      return res;
    }
    // ログインしていないユーザーをログインページにリダイレクト
    return NextResponse.redirect(new URL("/user/login", req.url));
  }

  if (req.nextUrl.pathname.startsWith("/api")) {
    if (handlingAccessToken()) {
      console.log("hi");
      return res;
    }
    console.log("gooo");

    return res;
  }
  return res;
}

export const config = {
  matcher: ["/cart/:path*", "/admin/:path*", "/api/:path*"],
};

const handlingAccessToken = (): boolean => {
  // ユーザーのログイン状態をチェックするロジックを追加
  const cookiesList = cookies();
  const hasAccessToken = cookiesList.has("access_token");
  if (hasAccessToken) {
    return true;
  }
  return false;
};
