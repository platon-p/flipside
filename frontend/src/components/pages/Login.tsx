import { useNavigate } from "react-router-dom"
import { Button } from "../shared/Button"
import { Input } from "../shared/Input"
import { useAuth } from "../../hooks/Auth"
import { useState } from "react";
import { AuthService } from "../../service/AuthService";

export function Login() {
  const auth = useAuth();
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("")

  function goToRegister() {
    navigate('/register');
  }

  function goToMain() {
    navigate('/')
  }

  async function submit() {
    const res = await auth.login(email, password);
    if (res) {
      setErrorMessage(res);
      return
    }
    goToMain();
  }

  if (auth.isAuth) {
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
      {errorMessage && <p style={{
        color: 'red'
      }}>{errorMessage}</p>}
      <p onClick={goToRegister}>Регистрация</p>
    </div>
  </>
}
