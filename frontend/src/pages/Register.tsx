import { Input } from "../components/Input"

function Register() {
    return <div style={{
        display: 'flex',
        flexDirection: 'column',
        gap: '0.2em'
    }}>
        <Input placeholder="Имя" />
        <Input placeholder="Никнейм" />
        <Input placeholder="Почта" />
        <Input placeholder="Пароль" type="password" />
        <Input placeholder="Повторите пароль" type="password" />

    </div>
}

export { Register }