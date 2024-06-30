import styles from "./index.module.scss";

export default function Header() {
  return (
    <header className={styles.header}>
      <div className={styles.site}>
        <a href="/">Skill Tracker</a>
      </div>
      <div className={styles.rightSide}>
        <ul className={styles.nav}>
          <li>
            <a href="/">スキル</a>
          </li>
          <li>
            <a href="/">プロフィール</a>
          </li>
          <li>
            <a href="/">ユーザ検索</a>
          </li>
          <li>
            <a href="/">管理メニュー</a>
            <ul className={styles.dropdown}>
              <li>
                <a href="/">スキル登録</a>
              </li>
              <li>
                <a href="/">ユーザ登録</a>
              </li>
            </ul>
          </li>
        </ul>
        <button>ログアウト</button>
      </div>
    </header>
  );
}
