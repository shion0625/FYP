import type { NextRequest } from "next/server";
import axios from "axios";
import { cookies } from "next/headers";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/login/`;
const COOKIE_SECURE =
  process.env.NEXT_PUBLIC_COOKIE_SECURE == "1" ? true : false;

export async function POST(req: NextRequest) {
  const json = await req.json();

  // GoのAPIを呼び出す
  const response = await axios.post(URL, json);

  // アクセストークンを取得
  const accessToken = response.headers["access_token"];

  // Cookieにアクセストークンを設定
  cookies().set({
    name: "access_token",
    value: accessToken,
    httpOnly: true,
    path: "/",
    sameSite: "strict",
    secure: COOKIE_SECURE,
  });

  // レスポンスを送信
  return Response.json(response.data);
}
