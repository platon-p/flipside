import { useNavigate } from "react-router-dom"
import { Button } from "../components/Button"
import { Input } from "../components/Input"

export function Register() {
    const navigate = useNavigate()

    function goToLogin() {
        navigate('/login')
    }

    function goToMain() {
        navigate('/')
    }

    return <div style={{
        display: 'flex',
        flexDirection: 'column',
        gap: '0.2em'
    }}>
        <h3 onClick={goToMain}>На главную</h3>
        <Input placeholder="Имя" />
        <Input placeholder="Никнейм" />
        <Input placeholder="Почта" />
        <Input placeholder="Пароль" type="password" />
        <Input placeholder="Повторите пароль" type="password" />
        <Button>Зарегистрироваться</Button>
        <p onClick={goToLogin}>Войти</p>
    </div>
}