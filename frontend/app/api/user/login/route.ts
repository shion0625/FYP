import type { NextRequest } from "next/server";
import { NextResponse } from "next/server";
import axios from "axios";
import { cookies } from "next/headers";

const URL = `${process.env.NEXT_PUBLIC_API_URL}/auth/login/`;

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
    sameSite: "strict"
  });

  // レスポンスを送信
  return Response.json(response.data);
}
