import { useReducer, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "@/hooks/Auth";
import { Button, Input } from "@/shared";

interface FormState {
  name: string;
  nickname: string;
  email: string;
  password: string;
  repeatPassword: string;
}

export function RegisterForm({ active }: { active: boolean }) {
  const navigate = useNavigate();
  const { register } = useAuth();

  const [formValues, handleInput] = useReducer(
    (
      state: FormState,
      action: [React.FormEvent<HTMLInputElement>, keyof FormState]
    ) => {
      return { ...state, [action[1]]: action[0].currentTarget.value };
    },
    {
      name: "",
      nickname: "",
      email: "",
      password: "",
      repeatPassword: "",
    }
  );
  const [errorRegMessage, setErrorMessage] = useState("");

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
      style={{ left: active ? 0 : "100%" }}
    >
      <Input placeholder="имя" onInput={(e) => handleInput([e, "name"])} />
      <Input
        placeholder="никнейм"
        onInput={(e) => handleInput([e, "nickname"])}
      />
      <Input placeholder="почта" onInput={(e) => handleInput([e, "email"])} />
      <Input
        placeholder="пароль"
        type="password"
        onInput={(e) => handleInput([e, "password"])}
      />
      <Input
        placeholder="повторите пароль"
        type="password"
        onInput={(e) => handleInput([e, "repeatPassword"])}
      />
      {errorRegMessage && <p className="text-red-500">{errorRegMessage}</p>}
      <Button onClick={submit}>зарегистрироваться</Button>
    </div>
  );
}
