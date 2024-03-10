import { useNavigate } from "react-router-dom"
import { Button } from "../shared/Button"
import { Input } from "../shared/Input"
import { useAuth } from "../../hooks/Auth";
import { useState } from "react";

export function Register() {
    const navigate = useNavigate();
    const { isAuth, register } = useAuth();
    const [formValues, setFormValues] = useState({
        name: '',
        nickname: '',
        email: '',
        password: '',
        repeatPassword: ''
    });
    const [errorMessage, setErrorMessage] = useState("");

    function goToLogin() {
        navigate('/login')
    }

    function goToMain() {
        navigate('/')
    }

    function handleInput(e: React.FormEvent<HTMLInputElement>, key: keyof typeof formValues) {
        setFormValues({
            ...formValues,
            [key]: e.currentTarget.value
        });
    }

    async function submit() {
        // validate
        if (formValues.password !== formValues.repeatPassword) {
            setErrorMessage('Пароли не совпадают');
            return;
        }
        // register
        const res = await register(formValues);
        if (res) {
            setErrorMessage(res);
            return;
        }
        goToMain();
    }

    if (isAuth) {
        goToMain();
    }

    return <div style={{
        display: 'flex',
        flexDirection: 'column',
        gap: '0.2em'
    }}>
        <h3 onClick={goToMain}>На главную</h3>
        <Input placeholder="Имя" onInput={e => handleInput(e, 'name')} />
        <Input placeholder="Никнейм" onInput={e => handleInput(e, 'nickname')} />
        <Input placeholder="Почта" onInput={e => handleInput(e, 'email')} />
        <Input placeholder="Пароль" type="password" onInput={e => handleInput(e, 'password')} />
        <Input placeholder="Повторите пароль" type="password" onInput={e => handleInput(e, 'repeatPassword')} />
        <Button onClick={submit}>Зарегистрироваться</Button>
        {errorMessage && <p style={{
            color: 'red'
        }}>{errorMessage}</p>}
        <p onClick={goToLogin}>Вход</p>
    </div>
}