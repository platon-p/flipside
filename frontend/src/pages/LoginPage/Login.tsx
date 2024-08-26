import { useState } from "react";
import { RegisterForm } from "./RegisterForm";
import { LoginForm } from "./LoginForm";
import { ButtonBox, ViewState } from "./ButtonBox";

export function LoginPage() {
  const [view, setView] = useState<ViewState>("login");

  return (
    <div className="px-8 pt-[20vh] max-w-md m-auto">
      <ButtonBox view={view} onChange={setView} />
      <div className="relative w-full h-72 overflow-hidden pt-1">
        <LoginForm active={view === "login"} />
        <RegisterForm active={view === "register"} />
      </div>
    </div>
  );
}

export default LoginPage;
