import { Button } from "../components/Button"
import { Input } from "../components/Input"

function Login() {
  return <>
  <div style={{
    display: 'flex',
    flexDirection: 'column',
    gap: '1em'
  }}>
    <Input placeholder="Почта" />
    <Input placeholder="Пароль" type="password" />
    <Button>Войти</Button>
  </div>
    
  </>
}

export {Login}
