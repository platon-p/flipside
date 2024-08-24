import { useAuth } from "@/hooks/Auth";
import { Button } from "@/shared/Button";
import { Input } from "@/shared/Input";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export function RegisterForm({ active }: { active: boolean }) {
  const navigate = useNavigate();
  const { register } = useAuth();
  const [formValues, setFormValues] = useState({
    name: "",
    nickname: "",
    email: "",
    password: "",
    repeatPassword: "",
  });
  const [errorRegMessage, setErrorRegMessage] = useState("");

  function handleInput(
    e: React.FormEvent<HTMLInputElement>,
    key: keyof typeof formValues
  ) {
    setFormValues({
      ...formValues,
      [key]: e.currentTarget.value,
    });
  }

  function goToMain() {
    navigate("/");
  }

  async function submitreg() {
    // validate
    if (formValues.password !== formValues.repeatPassword) {
      setErrorRegMessage("пароли не совпадают");
      return;
    }
    // register
    const res = await register(formValues);
    if (res) {
      setErrorRegMessage(res);
      return;
    }
    goToMain();
  }

  return (
    <div
      className="input-group"
      onSubmit={(e) => {
        e.preventDefault();
      }}
      style={{ left: active ? 0 : "100%" }}
    >
      <Input
        className="authorize"
        placeholder="имя"
        onInput={(e) => handleInput(e, "name")}
      />
      <Input
        placeholder="никнейм"
        onInput={(e) => handleInput(e, "nickname")}
        className="authorize"
      />
      <Input
        placeholder="почта"
        onInput={(e) => handleInput(e, "email")}
        className="authorize"
      />
      <Input
        placeholder="пароль"
        type="password"
        onInput={(e) => handleInput(e, "password")}
        className="authorize"
      />
      <Input
        placeholder="повторите пароль"
        type="password"
        onInput={(e) => handleInput(e, "repeatPassword")}
        className="authorize"
      />
      {errorRegMessage && (
        <p style={{ color: "red", fontFamily: "inter-norm" }}>
          {errorRegMessage}
        </p>
      )}
      <Button onClick={submitreg}>зарегистрироваться</Button>
    </div>
  );
}
