import { LoaderFunctionArgs, json, redirect } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

export async function loader({ request }: LoaderFunctionArgs) {
  const res = await fetch("http://localhost:8080/users/me", {
    method: "GET",
    headers: {
      Cookie: request.headers.get("Cookie") || "",
    },
    credentials: "include",
  });

  if (res.status !== 200) {
    return redirect("/login");
  }

  const resData = await res.json();

  return json({ user: resData });
}

export default function UserPage() {
  const { user } = useLoaderData<typeof loader>();

  return (
    <div>
      <h1>User Page</h1>
      <ul>
        <li>{user.email}</li>
        <li>
          {user.lastname} {user.firstname}
        </li>
        <li>{user.role}</li>
      </ul>
    </div>
  );
}
