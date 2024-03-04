import { useNavigate } from "react-router-dom"
import { Button } from "../components/Button"
import { Input } from "../components/Input"

export function Login() {
  const navigate = useNavigate()

  function goToRegister() {
    navigate('/register');
  }
  
  function goToMain() {
    navigate('/')
  }

  return <>
    <div style={{
      display: 'flex',
      flexDirection: 'column',
      gap: '0.2em'
    }}>
      <h3 onClick={goToMain}>На главную</h3>
      <Input placeholder="Почта" />
      <Input placeholder="Пароль" type="password" />
      <Button>Войти</Button>
      <p onClick={goToRegister}>Регистрация</p>
    </div>
  </>
}
