import { Link } from "@remix-run/react";
import styles from "./index.module.scss";

type Props = {
  isAdmin?: boolean;
};

export const Header: React.FC<Props> = ({ isAdmin = false }) => {
  return (
    <header className={styles.header}>
      <div className={styles.site}>
        <a href="/">Skill Tracker</a>
      </div>
      <div className={styles.rightSide}>
        <ul className={styles.nav}>
          <li>
            <a href="/user/skill/edit">スキル</a>
          </li>
          <li>
            <a href="/">プロフィール</a>
          </li>
          <li>
            <a href="/">ユーザ検索</a>
          </li>
          {isAdmin && (
            <li>
              <a href="/">管理メニュー</a>
              <ul className={styles.dropdown}>
                <li>
                  <a href="/admin/skill/create">スキル作成</a>
                </li>
                <li>
                  <a href="/">ユーザ登録</a>
                </li>
              </ul>
            </li>
          )}
        </ul>
        <Link to="/logout">ログアウト</Link>
      </div>
    </header>
  );
};

export default Header;
