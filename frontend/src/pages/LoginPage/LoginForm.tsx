import { useAuth } from "@/hooks/Auth";
import { Button } from "@/shared/Button";
import { Input } from "@/shared/Input";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

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
    const res = await auth.login(email, password);
    if (res) {
      setErrorLoginMessage(res);
      return;
    }
    goToMain();
  }

  return (
    <div className="input-group" style={{ left: active ? 0 : "-100%" }}>
      <Input
        className="authorize"
        onInput={(e) => setEmail(e.currentTarget.value)}
        placeholder="почта"
      />
      <Input
        className="authorize"
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
