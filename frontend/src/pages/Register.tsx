import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button, Input } from "@/shared";
import { useAuth } from "@/hooks/Auth";

export function Register() {
  const navigate = useNavigate();
  const { isAuth, register } = useAuth();
  const [formValues, setFormValues] = useState({
    name: "",
    nickname: "",
    email: "",
    password: "",
    repeatPassword: "",
  });
  const [errorMessage, setErrorMessage] = useState("");

  function goToLogin() {
    navigate("/login");
  }

  function goToMain() {
    navigate("/");
  }

  function handleInput(
    e: React.FormEvent<HTMLInputElement>,
    key: keyof typeof formValues
  ) {
    setFormValues({
      ...formValues,
      [key]: e.currentTarget.value,
    });
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

  if (isAuth) {
    goToMain();
  }

  return (
    <div className="flex flex-col gap-1">
      <h3 className="text-2xl" onClick={goToMain}>
        На главную
      </h3>
      <Input placeholder="Имя" onInput={(e) => handleInput(e, "name")} />
      <Input
        placeholder="Никнейм"
        onInput={(e) => handleInput(e, "nickname")}
      />
      <Input placeholder="Почта" onInput={(e) => handleInput(e, "email")} />
      <Input
        placeholder="Пароль"
        type="password"
        onInput={(e) => handleInput(e, "password")}
      />
      <Input
        placeholder="Повторите пароль"
        type="password"
        onInput={(e) => handleInput(e, "repeatPassword")}
      />
      <Button onClick={submit}>Зарегистрироваться</Button>
      {errorMessage && <p className="text-red-600">{errorMessage}</p>}
      <p onClick={goToLogin}>Вход</p>
    </div>
  );
}
