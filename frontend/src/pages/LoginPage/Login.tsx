import { useState } from "react";
import { RegisterForm } from "./RegisterForm";
import { LoginForm } from "./LoginForm";
import { ButtonBox, ViewState } from "./ButtonBox";

export default function Login() {
  const [view, setView] = useState<ViewState>("login");

  return (
    <div className="h-[70vh] mt-[20vh] max-w-md m-auto">
      <ButtonBox view={view} onChange={setView} />
      <div className="relative w-full h-full overflow-hidden pt-1">
        <LoginForm active={view === "login"} />
        <RegisterForm active={view === "register"} />
      </div>
    </div>
  );
}

