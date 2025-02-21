import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "@/hooks/Auth";
import { Button, Input } from "@/shared";
import { create, useStore } from "zustand";

interface FormState {
  name: string;
  nickname: string;
  email: string;
  password: string;
  repeatPassword: string;
}

const store = create<{
  formValues: FormState;
  setFormValues: (formValues: FormState) => void;
}>((set) => ({
  formValues: {
    name: "",
    nickname: "",
    email: "",
    password: "",
    repeatPassword: "",
  } as FormState,
  setFormValues: (formValues: FormState) => set({ formValues }),
}));

export function RegisterForm({ active }: { active: boolean }) {
  const navigate = useNavigate();
  const { register } = useAuth();

  const { formValues, setFormValues } = useStore(store);
  const handleInput = ([e, key]: [
    React.FormEvent<HTMLInputElement>,
    keyof FormState
  ]) => {
    setFormValues({
      ...formValues,
      [key]: e.currentTarget.value,
    });
  };
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
      className="w-full absolute flex flex-col gap-1 duration-300"
      style={{ left: active ? 0 : "130%" }}
    >
      <Input
        placeholder="имя"
        onInput={(e) => handleInput([e, "name"])}
      />
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
