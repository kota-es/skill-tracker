import React from "react";
import styles from "@/styles/routes/user.skill.register.module.scss";
import { Form, json, useLoaderData } from "@remix-run/react";
import { ActionFunctionArgs, LoaderFunctionArgs } from "@remix-run/node";
import Button from "@/components/shared/Button";
import { jsonWithError, jsonWithSuccess } from "remix-toast";

type SkillCategoryType = {
  id: number;
  name: string;
  skills: Skill[];
};

type Skill = {
  id: number;
  name: string;
  skill_category_id: number;
  description: string;
  level?: number;
  interested?: boolean;
};

type UserSkill = {
  id: number;
  user_id: number;
  skill_id: number;
  level: number;
  interested: boolean;
  created_at: string;
  updated_at: string;
};

type Request = {
  user_id: number;
  skills: SkillRequest[];
};

type SkillRequest = {
  skill_id: number;
  level: number;
  interested: boolean;
};

export const action = async ({ request }: ActionFunctionArgs) => {
  const formData: any = await request.formData();

  const sendData = {
    user_id: Number(formData.get("user_id")),
    skills: [],
  } as Request;
  formData.forEach((value: any, key: string) => {
    if (key === "user_id") return;

    const [type, strId] = key.split("_");
    const id = parseInt(strId);

    let skill = sendData.skills.find((skill) => skill.skill_id === id);
    if (!skill) {
      skill = { skill_id: id, level: 0, interested: false };
      sendData.skills.push(skill);
    }

    if (type === "level") {
      skill.level = Number(value);
    } else if (type === "interested") {
      skill.interested = true;
    }
  });

  const BASE_URL = import.meta.env.VITE_API_ORIGIN;
  const res = await fetch(`${BASE_URL}/users/skills`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(sendData),
  });

  if (res.ok) {
    return jsonWithSuccess({ result: null }, "スキルの登録が完了しました。");
  } else {
    return jsonWithError({ result: null }, "スキルの登録に失敗しました");
  }
};

export const loader = async ({ request }: LoaderFunctionArgs) => {
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

  let skills: Skill[] = [];
  if (SkillsRes.ok) {
    skills = await SkillsRes.json();
  }

  const userRes = await fetch(`${BASE_URL}/users/me`, {
    method: "GET",
    headers: {
      Cookie: request.headers.get("Cookie") || "",
    },
    credentials: "include",
  });
  const user = await userRes.json();

  const UserSkillRes = await fetch(`${BASE_URL}/users/${user.id}/skills`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  let userSkills: UserSkill[] = [];
  if (UserSkillRes.ok) {
    userSkills = await UserSkillRes.json();
  }

  const skillData = skillCategories.map((category) => {
    category.skills = skills.filter(
      (skill) => skill.skill_category_id === category.id
    );
    category.skills = category.skills.map((skill) => {
      const userSkill = userSkills.find(
        (userSkill) => userSkill.skill_id === skill.id
      );
      if (userSkill) {
        skill.level = userSkill.level;
        skill.interested = userSkill.interested;
      }
      return skill;
    });
    return category;
  });

  return json({ userId: user.id, skillData });
};

const SkillForm: React.FC = () => {
  const { userId, skillData: categories } = useLoaderData<typeof loader>();

  return (
    <>
      <div className={styles.container}>
        <h1>スキル編集</h1>
        <Form method="POST" className={styles.form}>
          {categories.map((category) => (
            <div key={category.id} className={styles.categorySection}>
              <h3 className={styles.categoryTitle}>{category.name}</h3>
              {category.skills.map((skill) => (
                <div key={skill.id} className={styles.skill}>
                  <label className={styles.skillLabel}>{skill.name}</label>
                  <div className={styles.levels}>
                    {[1, 2, 3, 4, 5].map((level) => (
                      <label key={level} className={styles.levelLabel}>
                        <input
                          type="radio"
                          name={`level_${skill.id}`}
                          value={level}
                          defaultChecked={skill.level === level}
                          className={styles.levelInput}
                        />
                        {level}
                      </label>
                    ))}
                  </div>
                  <label className={styles.interestLabel}>
                    <input
                      type="checkbox"
                      name={`interested_${skill.id}`}
                      value="true"
                      defaultChecked={skill.interested}
                      className={styles.interestInput}
                    />
                    関心あり
                  </label>
                </div>
              ))}
            </div>
          ))}
          <input type="hidden" name="user_id" value={userId} />
          <Button type="submit">登録</Button>
        </Form>
      </div>
    </>
  );
};

export default SkillForm;
