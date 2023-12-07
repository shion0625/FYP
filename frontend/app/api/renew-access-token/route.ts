import type { NextRequest } from "next/server";
import axios from "axios";
import { setAccessTokenCookie } from "@/utils/cookie";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/renew-access-token`;

export async function POST(req: NextRequest) {
  const json = await req.json();

  // GoのAPIを呼び出す
  const response = await axios.post(URL, json);

  // アクセストークンを取得
  const accessToken = response.data.data;
  // Cookieにアクセストークンを設定
  setAccessTokenCookie(accessToken);

  // レスポンスを送信
  return Response.json(response.data);
}
