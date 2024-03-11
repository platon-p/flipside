import { useNavigate } from "react-router-dom"
import { Button } from "@/components/shared/Button";
import './NewbiePage.css'

export function NewbiePage() {
    const navigate = useNavigate();
    function goToLogin() {
        navigate('/login')
    }

    return (
        <div>
            <div className="header">
                <div className="logo"></div>
                <div className="sign-in">
                    <a href="/login">войти</a>
                </div>
            </div>

            <div className="onboarding">
                <h1>flipside</h1>
                <div className="line"></div>
                <div className="info">
                    <p>Flipside - это сервис, благодоря которому
                        Вы можете значительно упростить процесс изучения
                        термиов и их понятий
                        <br />
                        <br></br>
                        Создайте свой первый набор карточек, выбрав тему, с которой вы хотели бы начать.
                        Это может быть что угодно - от языков и наук до искусства и спорта.
                    </p>
                </div>
                <Button className="create-set" onClick={goToLogin}>+ создать новый набор</Button>

            </div>
        </div>
    );
}