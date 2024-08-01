import Header from "@/components/Header";
import { User } from "@/types/User";
import { LoaderFunctionArgs } from "@remix-run/node";
import { Outlet, json, redirect, useLoaderData } from "@remix-run/react";

export async function loader({ request }: LoaderFunctionArgs) {
  const BASE_URL = import.meta.env.VITE_API_ORIGIN;

  const res = await fetch(`${BASE_URL}/users/me`, {
    method: "GET",
    headers: {
      Cookie: request.headers.get("Cookie") || "",
    },
    credentials: "include",
  });

  if (res.status !== 200) {
    return redirect("/login");
  }

  const resData: User = await res.json();

  return json({ user: resData });
}

export const Auth = () => {
  const { user } = useLoaderData<typeof loader>();

  return (
    <div>
      <Header isAdmin={user.role === "admin"} />
      <Outlet />
    </div>
  );
};

export default Auth;
