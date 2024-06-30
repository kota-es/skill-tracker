import React, { useState, useEffect } from "react";
import { Form } from "@remix-run/react";
import styles from "./index.module.scss";
import Header from "@/components/Header";
import Button from "@/components/shared/Button";
import { ActionFunctionArgs } from "@remix-run/node";
import { jsonWithSuccess } from "remix-toast";
interface SkillCategory {
  id: number;
  name: string;
}

export const action = async ({ request }: ActionFunctionArgs) => {
  const formData = await request.formData();
  const sendData = {
    name: formData.get("name"),
    description: formData.get("description"),
    is_new_category: formData.get("skill_category_option") === "new",
    skill_category_id: formData.get("skill_category_id"),
    skill_category_name: formData.get("skill_category_name"),
    level_explanation: [
      { level: 1, explanation: formData.get("level1") },
      { level: 2, explanation: formData.get("level2") },
      { level: 3, explanation: formData.get("level3") },
      { level: 4, explanation: formData.get("level4") },
      { level: 5, explanation: formData.get("level5") },
    ],
  };

  const BASE_URL = import.meta.env.VITE_API_ORIGIN;

  const res = await fetch(`${BASE_URL}/admin/skills`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(sendData),
  });

  if (res.status === 200) {
    return { status: 200, headers: { Location: "/admin/skills" } };
  }

  return jsonWithSuccess(
    { result: "Data saved successfully" },
    "Operation successful! 🎉"
  );
};

const SkillCreate: React.FC = () => {
  const [skillCategories, setSkillCategories] = useState<SkillCategory[]>([]);
  const [categoryOption, setCategoryOption] = useState<"existing" | "new">(
    "existing"
  );

  const handleCategoryOptionChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setCategoryOption(e.target.value as "existing" | "new");
  };

  return (
    <>
      <Header />
      <div className={styles.container}>
        <h1>スキル作成</h1>
        <Form method="POST" className={styles.form}>
          <div className={styles.formGroup}>
            <label>スキル名:</label>
            <input type="text" name="name" required />
          </div>

          <div className={styles.formGroup}>
            <label>スキルカテゴリ:</label>
            <div className={styles.categoryOptions}>
              <label>
                <input
                  type="radio"
                  value="existing"
                  name="skill_category_option"
                  checked={categoryOption === "existing"}
                  onChange={handleCategoryOptionChange}
                />
                既存のカテゴリを選択
              </label>
              <label>
                <input
                  type="radio"
                  value="new"
                  name="skill_category_option"
                  checked={categoryOption === "new"}
                  onChange={handleCategoryOptionChange}
                />
                新しいカテゴリを作成
              </label>
            </div>
            {categoryOption === "existing" && (
              <select name="skill_category_id ">
                <option value="">カテゴリを選択</option>
                {skillCategories.map((category) => (
                  <option key={category.id} value={category.id}>
                    {category.name}
                  </option>
                ))}
              </select>
            )}
            {categoryOption === "new" && (
              <input
                type="text"
                name="skill_category_name"
                placeholder="カテゴリ名を入力"
                required
              />
            )}
          </div>

          <div className={styles.formGroup}>
            <label>スキル説明:</label>
            <textarea name="description" />
          </div>

          <div className={styles.formGroup}>
            <label>レベルごとの説明:</label>
            {[1, 2, 3, 4, 5].map((level, index) => (
              <div key={index} className="level-description">
                <label>レベル {index + 1}:</label>
                <textarea name={`level${level}`} required />
              </div>
            ))}
          </div>

          <Button>作成</Button>
        </Form>
      </div>
    </>
  );
};

export default SkillCreate;
