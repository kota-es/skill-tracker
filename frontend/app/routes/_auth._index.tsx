import { User } from "@/types/User";
import { useRouteLoaderData } from "@remix-run/react";

type RouteData = {
  user: User;
};

export default function UserPage() {
  const { user } = useRouteLoaderData("routes/_auth") as RouteData;

  return (
    <div>
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
