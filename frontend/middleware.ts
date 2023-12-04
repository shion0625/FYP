import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { NextPageContext } from "next";

export function middleware(req: NextRequest) {
  // ユーザーのログイン状態をチェックするロジックを追加

  if (
    req.nextUrl.pathname.startsWith("/cart") ||
    req.nextUrl.pathname.startsWith("/admin")
  ) {
    // ログインしていないユーザーをログインページにリダイレクト
    return NextResponse.redirect(new URL("/user/login", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/cart/:path*", "/admin/:path*"],
};
