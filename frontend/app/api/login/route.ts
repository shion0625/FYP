import type { NextRequest } from "next/server";
import axios from "axios";
import { setAccessTokenCookie, setRefreshTokenCookie } from "@/utils/cookie";
const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/login/`;

export async function POST(req: NextRequest) {
  const json = await req.json();

  // GoのAPIを呼び出す
  const response = await axios.post(URL, json);

  // アクセストークンを取得
  const { accessToken, refreshToken } = response.data.data;

  // Cookieにアクセストークンを設定
  setAccessTokenCookie(accessToken);
  setRefreshTokenCookie(refreshToken);
  // レスポンスを送信
  return Response.json(response.data);
}
