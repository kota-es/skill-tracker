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

  const user: User = await res.json();

  if (user.role !== "admin") {
    return redirect("/");
  }

  return json({ user });
}

export const Admin = () => {
  const { user } = useLoaderData<typeof loader>();

  return (
    <div>
      <Header isAdmin={true} />
      <Outlet />
    </div>
  );
};

export default Admin;
