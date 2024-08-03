import Button from "@/components/shared/Button";
import { Form } from "@remix-run/react";

import type { ProfileData } from "@/types/ProfileData";

import styles from "./index.module.scss";

type Props = {
  profileData: ProfileData;
  userId: number;
};

export const UserProfileEditPage: React.FC<Props> = ({
  profileData,
  userId,
}) => {
  return (
    <>
      <div className={styles.container}>
        <h1>プロフィール編集</h1>
        <Form method="POST" className={styles.form}>
          <div className={styles.formGroup}>
            <label>経験：</label>
            <textarea name="notes" rows={5} defaultValue={profileData.notes} />
          </div>
          <div className={styles.formGroup}>
            <label>やりたいこと:</label>
            <textarea
              name="desires"
              rows={5}
              defaultValue={profileData.desires}
            />
          </div>
          <div className={styles.formGroup}>
            <label>やりたくないこと:</label>
            <textarea
              name="dislikes"
              rows={5}
              defaultValue={profileData.dislikes}
            />
          </div>
          <input type="hidden" name="user_id" value={userId} />
          <Button type="submit">登録</Button>
        </Form>
      </div>
    </>
  );
};
