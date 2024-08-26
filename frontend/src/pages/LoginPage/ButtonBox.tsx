export type ViewState = "login" | "register";

export interface ButtonBoxProps {
  view: ViewState;
  onChange: (view: ViewState) => void;
}

export function ButtonBox({ view, onChange }: ButtonBoxProps) {
  return (
    <div>
      <div>
        <button
          className="w-1/2 h-8"
          onClick={() => onChange("login")}
          style={{ color: view === "login" ? "#F1694F" : undefined }}
        >
          вход
        </button>
        <button
          className="w-1/2 h-8"
          onClick={() => onChange("register")}
          style={{ color: view === "register" ? "#F1694F" : undefined }}
        >
          регистрация
        </button>
      </div>
      <div className="relative w-full h-1">
        <div className="absolute w-full h-full bg-[#D9D9D9] duration-150" />
        <div
          className="absolute h-full bg-orange-500 duration-150 w-1/2"
          style={{ left: view === "login" ? 0 : "50%" }}
        />
      </div>
    </div>
  );
}
