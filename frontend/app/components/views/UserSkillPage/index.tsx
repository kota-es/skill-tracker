import Header from "@/components/Header";
import type { UserSkillData } from "@/types/UserSkillData";
import styles from "./index.module.scss";

type Props = {
  skillData: UserSkillData;
};

const UserSkillPage: React.FC<Props> = ({ skillData }) => {
  return (
    <div>
      <Header />
      <div className={styles.skillListContainer}>
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
    </div>
  );
};

export default UserSkillPage;
