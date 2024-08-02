import { createCookie } from "@remix-run/node";
import { redirectWithToast } from "remix-toast";

const accessTokenCookie = createCookie("access_token");

export const loader = async () => {
  const headers = new Headers();
  headers.append(
    "Set-Cookie",
    await accessTokenCookie.serialize("/", {
      expires: new Date(0),
    })
  );

  return redirectWithToast(
    "/login",
    { message: "ログアウトしました", type: "success" },
    { headers }
  );
};
