import { useNavigate } from "react-router-dom"
import { Button } from "../components/Button"
import { Input } from "../components/Input"
import { useAuth } from "../hooks/Auth"
import { useState } from "react";
import { AuthService } from "../service/AuthService";

export function Login() {
  const { isAuth } = useAuth();
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  function goToRegister() {
    navigate('/register');
  }

  function goToMain() {
    navigate('/')
  }

  async function submit() {
    const res = await AuthService.loginByEmail(email, password);
    console.log(res);
  }

  if (isAuth) {
    goToMain();
  }

  return <>
    <div style={{
      display: 'flex',
      flexDirection: 'column',
      gap: '0.2em'
    }}>
      <h3 onClick={goToMain}>На главную</h3>
      <Input onInput={(event) => {
        setEmail(event.currentTarget.value)
      }} placeholder="Почта" />
      <Input onInput={(event) => {
        setPassword(event.currentTarget.value)
      }} placeholder="Пароль" type="password" />
      <Button onClick={submit}>Войти</Button>
      <p onClick={goToRegister}>Регистрация</p>
    </div>
  </>
}
