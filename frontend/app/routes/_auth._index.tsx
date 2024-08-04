import { UserListPage } from "@/components/views/UserListPage";
import { SearchedUser } from "@/types/UserList";
import { json, useLoaderData } from "@remix-run/react";

export const loader = async () => {
  const BASE_URL = import.meta.env.VITE_API_ORIGIN;

  const res = await fetch(`${BASE_URL}/users/search`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (res.ok) {
    const userList: SearchedUser[] = await res.json();
    const formatUserList: SearchedUser[] = userList.map((user) => {
      const formatedProfile = {
        ...user.profile,
        created_at: formatDate(user.profile.created_at),
        updated_at: formatDate(user.profile.updated_at),
      };

      const formatedSkills = user.skills.map((skill) => {
        return {
          ...skill,
          created_at: formatDate(skill.created_at),
          updated_at: formatDate(skill.updated_at),
        };
      });

      const lastUpdatedSkill = formatedSkills.reduce((acc, skill) => {
        if (acc === null) {
          return skill;
        }

        if (new Date(acc.updated_at) < new Date(skill.updated_at)) {
          return skill;
        }

        return acc;
      });

      return {
        ...user,
        profile: formatedProfile,
        skills: formatedSkills,
        lastUpdatedSkillDate: lastUpdatedSkill.updated_at,
      };
    });

    return json({ userList: formatUserList });
  }

  return null;
};

const formatDate = (dateStr: string) => {
  // FIXME: APIの仕様変更後に削除
  if (dateStr === "0001-01-01T09:18:59+09:18") {
    return "";
  }
  const date = new Date(dateStr);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const day = String(date.getDate()).padStart(2, "0");
  const hours = String(date.getHours()).padStart(2, "0");
  const minutes = String(date.getMinutes()).padStart(2, "0");
  const seconds = String(date.getSeconds()).padStart(2, "0");

  return `${year}/${month}/${day} ${hours}:${minutes}:${seconds}`;
};

export const UserPage = () => {
  const { userList } = useLoaderData<typeof loader>();

  return <UserListPage userList={userList} />;
};

export default UserPage;
