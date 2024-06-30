import React from "react";
import styles from "./index.module.scss";

type ButtonProps = {
  type?: "button" | "submit";
  variant?:
    | "primary"
    | "secondary"
    | "success"
    | "danger"
    | "warning"
    | "info"
    | "light"
    | "dark";
  onClick?: () => void;
  children?: React.ReactNode;
};

const Button: React.FC<ButtonProps> = ({
  type = "submit",
  variant = "primary",
  onClick,
  children,
}) => {
  return (
    <button
      type={type}
      className={`${styles.button} ${styles[`button--${variant}`]}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

export default Button;
