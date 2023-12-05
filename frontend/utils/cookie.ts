import { cookies } from "next/headers";

const COOKIE_SECURE =
  process.env.NEXT_PUBLIC_COOKIE_SECURE == "1" ? true : false;
const accessTokenExpiresInMinutes = 20;
const refreshTokenExpiresInDays = 7;

export function setAccessTokenCookie(accessToken: string) {
  const expiryDate = new Date();
  expiryDate.setMinutes(expiryDate.getMinutes() + accessTokenExpiresInMinutes);

  cookies().set({
    name: "access_token",
    value: accessToken,
    httpOnly: true,
    path: "/",
    sameSite: "strict",
    secure: COOKIE_SECURE,
    expires: expiryDate,
  });
}

export function setRefreshTokenCookie(refreshToken: string) {
  const expiryDate = new Date();
  expiryDate.setDate(expiryDate.getDate() + refreshTokenExpiresInDays);
  cookies().set({
    name: "refresh_token",
    value: refreshToken,
    httpOnly: true,
    path: "/",
    sameSite: "strict",
    secure: COOKIE_SECURE,
    expires: expiryDate,
  });
}
