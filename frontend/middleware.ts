import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { cookies } from "next/headers";

export function middleware(req: NextRequest) {
  const response = NextResponse.next();

  // ユーザーのログイン状態をチェックするロジックを追加
  const cookiesList = cookies();
  const hasAccessToken = cookiesList.has("access_token");
  if (
    (!hasAccessToken && req.nextUrl.pathname.startsWith("/cart")) ||
    req.nextUrl.pathname.startsWith("/admin")
  ) {
    // ログインしていないユーザーをログインページにリダイレクト
    return NextResponse.redirect(new URL("/user/login", req.url));
  }

  return response;
}

export const config = {
  matcher: ["/cart/:path*", "/admin/:path*", "/api/:path*"],
};
