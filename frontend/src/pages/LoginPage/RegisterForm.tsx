import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "@/hooks/Auth";
import { Button, Input } from "@/shared";

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
  const [errorRegMessage, setErrorMessage] = useState("");

  function handleInput(
    e: React.FormEvent<HTMLInputElement>,
    key: keyof typeof formValues,
  ) {
    setFormValues({
      ...formValues,
      [key]: e.currentTarget.value,
    });
  }

  function goToMain() {
    navigate("/");
  }

  function submit() {
    // validate
    if (formValues.password !== formValues.repeatPassword) {
      setErrorMessage("пароли не совпадают");
      return;
    }
    // register
    register(formValues)
      .then(() => {
        goToMain();
      })
      .catch((e: Error) => {
        setErrorMessage(e.message);
      });
  }

  return (
    <div
      className="w-full absolute flex flex-col gap-1 duration-200"
      onSubmit={(e) => {
        e.preventDefault();
      }}
      style={{ left: active ? 0 : "100%" }}
    >
      <Input placeholder="имя" onInput={(e) => handleInput(e, "name")} />
      <Input
        placeholder="никнейм"
        onInput={(e) => handleInput(e, "nickname")}
      />
      <Input placeholder="почта" onInput={(e) => handleInput(e, "email")} />
      <Input
        placeholder="пароль"
        type="password"
        onInput={(e) => handleInput(e, "password")}
      />
      <Input
        placeholder="повторите пароль"
        type="password"
        onInput={(e) => handleInput(e, "repeatPassword")}
      />
      {errorRegMessage && (
        <p style={{ color: "red", fontFamily: "inter-norm" }}>
          {errorRegMessage}
        </p>
      )}
      <Button onClick={submit}>зарегистрироваться</Button>
    </div>
  );
}
