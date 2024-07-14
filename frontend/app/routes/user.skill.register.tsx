import React, { useState } from "react";
import styles from "@/styles/routes/user.skill.register.module.scss";
import Header from "@/components/Header";
import { Form, json, useLoaderData } from "@remix-run/react";
import { ActionFunctionArgs } from "@remix-run/node";
import Button from "@/components/shared/Button";

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

  const sendData = { user_id: 1, skills: [] } as Request;
  formData.forEach((value: any, key: string) => {
    const [type, strId] = key.split("_");
    const id = parseInt(strId);

    let skill = sendData.skills.find((skill) => skill.skill_id === id);
    if (!skill) {
      skill = { skill_id: id, level: 0, interested: false };
      sendData.skills.push(skill);
    }

    if (type === "level") {
      skill.level = value;
    } else if (type === "interested") {
      skill.interested = true;
    }
  });

  console.log("Send Data:", sendData);
  return json({ success: true });
};

export const loader = async () => {
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

  const skillData = skillCategories.map((category) => {
    category.skills = skills.filter(
      (skill) => skill.skill_category_id === category.id
    );
    return category;
  });

  return json({ skillData });
};

const SkillForm: React.FC = () => {
  const [skillLevels, setSkillLevels] = useState<{ [key: string]: number }>({});
  const { skillData: categories } = useLoaderData<typeof loader>();

  const handleLevelChange = (skillName: string, level: number) => {
    setSkillLevels({ ...skillLevels, [skillName]: level });
  };

  return (
    <>
      <Header />
      <div className={styles.container}>
        <h1>スキル登録</h1>
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
                          checked={skillLevels[skill.name] === level}
                          onChange={() => handleLevelChange(skill.name, level)}
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
                      className={styles.interestInput}
                    />
                    関心あり
                  </label>
                </div>
              ))}
            </div>
          ))}
          <Button type="submit">登録</Button>
        </Form>
      </div>
    </>
  );
};

export default SkillForm;
