import Header from "@/components/Header";
import type { UserSkillData } from "@/types/UserSkillData";
import styles from "./index.module.scss";
import { ProfileData } from "@/types/ProfileData";
import { User } from "@/types/User";

type Props = {
  skillData: UserSkillData;
  profileData: ProfileData;
  user: User;
};

const UserSkillPage: React.FC<Props> = ({ skillData, profileData, user }) => {
  return (
    <>
      <div className={styles.skillListContainer}>
        <h2>
          {user.lastname} {user.firstname}
        </h2>
        <div className={styles.profileSection}>
          <h3>経験</h3>
          <p>{profileData.notes || "未登録"}</p>
          <h3>やりたいこと</h3>
          <p>{profileData.desires || "未登録"}</p>
          <h3>やりたくないこと</h3>
          <p>{profileData.dislikes || "未登録"}</p>
        </div>
        {skillData.map((category) => (
          <div key={category.id} className={styles.categorySection}>
            <h3 className={styles.categoryName}>{category.name}</h3>
            <div className={styles.table}>
              {category.skills.map((skill) => (
                <div key={skill.id} className={styles.tableRow}>
                  <div className={styles.tableCell}>
                    <span>{skill.name}</span>
                    {skill.interested && (
                      <span className={styles.interestedBadge}>興味あり</span>
                    )}
                  </div>
                  <span
                    className={`${styles.tableCell} ${styles.skillLevel} ${
                      styles[`level${skill.level}`]
                    }`}
                  >
                    Level: {skill.level}
                  </span>
                  <span className={styles.tableCellWide}>
                    {skill.level > 0 ? skill.levelExplanation : "未経験"}
                  </span>
                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
    </>
  );
};

export default UserSkillPage;
