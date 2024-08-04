import { SearchedUser } from "@/types/UserList";
import styles from "./index.module.scss";
import { useNavigate } from "@remix-run/react";

type Props = {
  userList: SearchedUser[];
};

export const UserListPage: React.FC<Props> = ({ userList }) => {
  const navigate = useNavigate();
  console.log(userList);
  return (
    <>
      <div className={styles.container}>
        <h1>ユーザー一覧</h1>
        <table className={styles.table}>
          <thead>
            <tr>
              <th>氏名(かな)</th>
              <th>種別</th>
              <th>プロフィール更新日</th>
              <th>スキル更新日</th>
            </tr>
          </thead>
          <tbody>
            {userList.map((user) => (
              <tr
                key={user.user_id}
                onClick={() => navigate(`user/${user.user_id}`)}
              >
                <td>
                  {user.lastname} {user.firstname}（{user.lastname_kana}{" "}
                  {user.firstname_kana}）
                </td>
                <td>{user.role === "admin" ? "管理" : "一般"}</td>
                <td>{user.profile.updated_at || "未設定"}</td>
                <td>{user.lastUpdatedSkillDate || "未設定"}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
};
