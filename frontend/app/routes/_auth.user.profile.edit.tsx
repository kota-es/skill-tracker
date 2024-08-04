import React from "react";
import { json, useLoaderData } from "@remix-run/react";
import { ActionFunctionArgs, LoaderFunctionArgs } from "@remix-run/node";
import { UserProfileEditPage } from "@/components/views/UserProfileEditPage";

import type { ProfileData } from "@/types/ProfileData";
import { jsonWithError, jsonWithSuccess } from "remix-toast";

export const action = async ({ request }: ActionFunctionArgs) => {
  const formData: any = await request.formData();
  const sendData = {
    user_id: Number(formData.get("user_id")),
    notes: formData.get("notes"),
    desires: formData.get("desires"),
    dislikes: formData.get("dislikes"),
  };

  const BASE_URL = import.meta.env.VITE_API_ORIGIN;

  const res = await fetch(`${BASE_URL}/users/profile`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(sendData),
  });

  if (res.ok) {
    return jsonWithSuccess({}, "プロフィールの編集が完了しました。");
  } else {
    return jsonWithError({}, "プロフィールの編集に失敗しました");
  }
};

export const loader = async ({ request }: LoaderFunctionArgs) => {
  const BASE_URL = import.meta.env.VITE_API_ORIGIN;

  const userRes = await fetch(`${BASE_URL}/users/me`, {
    method: "GET",
    headers: {
      Cookie: request.headers.get("Cookie") || "",
    },
    credentials: "include",
  });
  const user = await userRes.json();

  const ProfileRes = await fetch(`${BASE_URL}/users/${user.id}/profile`, {
    method: "GET",
    credentials: "include",
  });
  const profileData: ProfileData = await ProfileRes.json();
  console;

  return json({ userId: user.id, profileData });
};

export const UserProfileEdit: React.FC = () => {
  const { userId, profileData } = useLoaderData<typeof loader>();

  return <UserProfileEditPage profileData={profileData} userId={userId} />;
};

export default UserProfileEdit;
