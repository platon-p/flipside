import { useNavigate } from "react-router-dom"
import { useState } from 'react'
import Navbar from "../../Navbar";
import './NewbiePage.css'
export function NewbiePage() {
    const [request, setRequest] = useState('');

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
                    <h3>Flipside - это разъёб <br/>
                    ...<br/>
                    ...<br/>
                    ...<br/>
                    ...<br/>
                    ...<br/>
                    ...<br/>
                    ...<br/>
                    ...<br/>
                    ...<br/>
                    ...<br/>
                    ...<br/>
                    </h3>
                </div>
                
            </div>
		</div>
	);
}