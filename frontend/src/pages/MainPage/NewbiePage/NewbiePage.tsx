import { useNavigate } from "react-router-dom";
import { OutlineButton } from "@/shared/OutlineButton/OutlineButton";
import { Header } from "@/widgets/Header/Header";

export function NewbiePage() {
  const navigate = useNavigate();
  function goToLogin() {
    navigate("/login");
  }

  return (
    <div className="px-2">
      <Header />
      <div className="flex w-auto flex-wrap">
        <h1 className="text-3xl">flipside</h1>
        <div className="w-full h-0.5 bg-[#D9D9D9]" />
        <div>
          <p className="text-left">
            Flipside - это сервис, благодоря которому Вы можете значительно
            упростить процесс изучения термиов и их понятий
            <br />
            <br></br>
            Создайте свой первый набор карточек, выбрав тему, с которой вы
            хотели бы начать. Это может быть что угодно - от языков и наук до
            искусства и спорта.
          </p>
        </div>
        <OutlineButton className="w-full" onClick={goToLogin}>
          создать новый набор
        </OutlineButton>
      </div>
    </div>
  );
}
