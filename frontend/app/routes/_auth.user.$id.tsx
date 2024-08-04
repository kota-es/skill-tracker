import { json, useLoaderData, useParams } from "@remix-run/react";
import UserSkillPage from "@/components/views/UserSkillPage";
import type { UserSkillData, UserSKills } from "@/types/UserSkillData";
import { LoaderFunctionArgs } from "@remix-run/node";

interface SkillCategoryType {
  id: number;
  name: string;
}

type UserSkill = {
  id: number;
  user_id: number;
  skill_id: number;
  level: number;
  interested: boolean;
  created_at: string;
  updated_at: string;
};

type SkillData = {
  id: number;
  name: string;
  skill_category_id: number;
  description: string;
  levels: SkillLevel[];
};

type SkillLevel = {
  level: number;
  explanation: string;
};

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const BASE_URL = import.meta.env.VITE_API_ORIGIN;

  const categoryRes = await fetch(`${BASE_URL}/skills/categories`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  let skillCategories: SkillCategoryType[] = [];
  if (categoryRes.ok) {
    skillCategories = await categoryRes.json();
  }

  const SkillsRes = await fetch(`${BASE_URL}/skills`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  let skillData: SkillData[] = [];
  if (SkillsRes.ok) {
    skillData = await SkillsRes.json();
  }

  const userSkillRes = await fetch(`${BASE_URL}/users/${params.id}/skills`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  let userSkills: any = [];
  if (userSkillRes.ok) {
    userSkills = await userSkillRes.json();
  }

  const userSkillData: UserSkillData = skillCategories.map(
    (category: SkillCategoryType) => {
      const skillsInCategory = skillData.filter(
        (skill: SkillData) => skill.skill_category_id === category.id
      );
      const userSkillsInCategory = userSkills.filter((userSkill: UserSkill) => {
        return skillsInCategory.some(
          (skill) => skill.id === userSkill.skill_id
        );
      });

      const skillsData: UserSKills[] = skillsInCategory.map(
        (skill: SkillData) => {
          const userSkill = userSkillsInCategory.find(
            (userSkill: UserSkill) => userSkill.skill_id === skill.id
          );

          const skillLevel = skill.levels.find(
            (level: SkillLevel) => level.level === userSkill?.level
          );

          return {
            id: skill.id,
            name: skill.name,
            description: skill.description,
            level: userSkill?.level || 0,
            levelExplanation: skillLevel?.explanation || "",
            interested: userSkill?.interested || false,
          };
        }
      );

      return {
        id: category.id,
        name: category.name,
        skills: skillsData,
      };
    }
  );

  return json({ userSkillData });
};

export const UserSkill = () => {
  const { userSkillData } = useLoaderData<typeof loader>();

  return <UserSkillPage skillData={userSkillData} />;
};

export default UserSkill;
