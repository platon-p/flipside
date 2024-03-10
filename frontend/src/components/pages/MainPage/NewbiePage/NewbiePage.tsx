import { useNavigate } from "react-router-dom"
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
                    <a onClick={goToLogin}>войти</a>
                </div>
            </div>

            <div className="onboarding">
                <h1>flipside</h1>
                <div className="line"></div>
                <div className="info">
                    <h3>Flipside - это круто </h3>
                </div>
            </div>
        </div>
    );
}