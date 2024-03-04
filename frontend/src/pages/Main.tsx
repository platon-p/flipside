import { useNavigate } from "react-router-dom"

export function Main() {
    const navigate = useNavigate()

    function goToLogin() {
        navigate('/login')
    }
    function goToRegister() {
        navigate('/register')
    }
    return <>
        <h1>Index</h1 >
        <p onClick={goToLogin}>Вход</p>
        <p onClick={goToRegister}>Регистрация</p>
    </>
}