// Filename - "./components/Navbar.js
import { useLocation } from "react-router-dom";

const Navbar = () => {
	const location = useLocation();
	if (location.pathname != '/authorize'){
	return (
		<div className="navbar">
			<ul className="navbar-menu">
				<li><a href ="/">главная</a></li>	
				<li><a href="/authorize">создать</a></li>
				<li><a href="/login">профиль</a></li>
			</ul>
		</div>
	);}
};

export default Navbar;
