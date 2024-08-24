import { useNavigate } from "react-router-dom";
import css from "./NewbiePage.module.css";
import { Button } from "@/shared/Button/Button";
import { Header } from "@/shared/Header/Header";

export function NewbiePage() {
  const navigate = useNavigate();
  function goToLogin() {
    navigate("/login");
  }

  return (
    <div style={{padding: '0 1em'}}>
      <Header />
      <div className={css.onboarding}>
        <h1>flipside</h1>
        <div
          style={{
            width: "100%",
            backgroundColor: "rgba(217, 217, 217, 1)",
            height: "2px",
          }}
        ></div>
        <div>
          <p style={{ textAlign: "left" }}>
            Flipside - это сервис, благодоря которому Вы можете значительно
            упростить процесс изучения термиов и их понятий
            <br />
            <br></br>
            Создайте свой первый набор карточек, выбрав тему, с которой вы
            хотели бы начать. Это может быть что угодно - от языков и наук до
            искусства и спорта.
          </p>
        </div>
        <Button
          style={{
            width: "100%",
          }}
          onClick={goToLogin}
        >
          создать новый набор
        </Button>
      </div>
    </div>
  );
}
