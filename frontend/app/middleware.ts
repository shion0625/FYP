import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export function middleware(request: NextRequest) {
  // ユーザーのログイン状態をチェックするロジックを追加
  const isLoggedIn = checkUserLoginStatus(request);

  if (
    !isLoggedIn &&
    (request.nextUrl.pathname.startsWith("/cart") ||
      request.nextUrl.pathname.startsWith("/admin"))
  ) {
    // ログインしていないユーザーをログインページにリダイレクト
    return NextResponse.redirect("/login");
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/cart/*", "/admin/*"],
};

// ユーザーのログイン状態をチェックする関数
function checkUserLoginStatus(request: NextRequest) {
  // ここにログイン状態をチェックするロジックを実装
  // 例えば、Cookieやセッションをチェックするなど
  return true;
}
