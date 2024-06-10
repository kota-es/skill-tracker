import { ActionFunctionArgs, json, redirect } from "@remix-run/node";
import { Form } from "@remix-run/react";
import styles from "./index.module.scss";

export async function action({ request }: ActionFunctionArgs) {
  const formData = await request.formData();

  const sendData = {
    email: formData.get("email"),
    password: formData.get("password"),
  };

  const res = await fetch("http://localhost:8080/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(sendData),
  });

  if (res.status === 200) {
    const cookieData = res.headers.get("Set-Cookie");

    if (!cookieData) {
      return json({ error: "ログインに失敗しました" }, { status: 500 });
    }

    const cookies = Object.fromEntries(
      cookieData.split(", ").map((cookie) => {
        const [keyValue, ...options] = cookie.split("; ");
        const [key, value] = keyValue.split("=");
        return [key, { value, options }];
      })
    );

    const accessToken = cookies["access_token"].value;

    const headers = new Headers();
    headers.append(
      "Set-Cookie",
      "access_token=" +
        accessToken +
        "; Max-Age=900; HttpOnly; Path=/" +
        "; Secure"
    );

    return redirect("/user", { headers });
  }

  return null;
}

export default function LoginPage() {
  return (
    <div className={styles.loginBg}>
      <div className={styles.loginContainer}>
        <h2>Skill Tracker</h2>
        <Form method="POST" className={styles.loginForm}>
          <div className={styles.inputGroup}>
            <label htmlFor="username">メールアドレス</label>
            <input type="email" id="email" name="email" />
          </div>
          <div className={styles.inputGroup}>
            <label htmlFor="password">パスワード</label>
            <input type="password" id="password" name="password" />
          </div>
          <button type="submit">ログイン</button>
        </Form>
      </div>
    </div>
  );
}
