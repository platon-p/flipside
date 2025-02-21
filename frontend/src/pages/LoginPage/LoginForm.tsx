import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "@/hooks/Auth";
import { Button, Input } from "@/shared";

export function LoginForm({ active }: { active: boolean }) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorLoginMessage, setErrorLoginMessage] = useState("");
  const auth = useAuth();
  const navigate = useNavigate();

  function goToMain() {
    navigate("/");
  }

  async function submit() {
    auth
      .login(email, password)
      .then(() => {
        goToMain();
      })
      .catch((e: Error) => {
        setErrorLoginMessage(e.message);
      });
  }

  return (
    <div
      className="w-full absolute flex flex-col gap-1 duration-300" 
      style={{ left: active ? 0 : "-130%" }}
    >
      <Input
        onInput={(e) => setEmail(e.currentTarget.value)}
        placeholder="почта"
      />
      <Input
        onInput={(e) => setPassword(e.currentTarget.value)}
        placeholder="пароль"
        type="password"
      />
      {errorLoginMessage && (
        <p style={{ color: "red", fontFamily: "inter-norm" }}>
          {errorLoginMessage}
        </p>
      )}
      <Button onClick={submit}>войти</Button>
    </div>
  );
}
