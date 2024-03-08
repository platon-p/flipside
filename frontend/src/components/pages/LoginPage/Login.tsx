import { useNavigate } from "react-router-dom"
import { Button } from "@/components/shared/Button"
import { Input } from "@/components/shared/Input"
import { useAuth } from "@/hooks/Auth"
import { useState } from "react";
import './Login.css';

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
      <h3>Вход</h3>
      <div className="login-form">
        <Input onInput={e => setEmail(e.currentTarget.value)} placeholder="Почта" />
        <Input onInput={e => setPassword(e.currentTarget.value)} placeholder="Пароль" type="password" />
        <Button onClick={submit}>Войти</Button>
        {errorMessage && <p className='error-message'>{errorMessage}</p>}
      </div>
      <div>
        <Button onClick={goToRegister}>К регистрации</Button>
        <Button onClick={goToMain}>На главную</Button>
      </div>
    </div>
  </>
}
