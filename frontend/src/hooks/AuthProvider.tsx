import { useEffect, useState } from "react";
import { AuthData, RegisterData, authContext } from "@/hooks/Auth";
import { AuthService } from "@/service/AuthService";

const emailPattern = /^\s*[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]+\s*$/;

const useProvideAuth = (): AuthData => {
  const [isAuth, setIsAuth] = useState<boolean>(AuthService.isAuth());

  function logout() {
    AuthService.logout();
    setIsAuth(false);
  }

  async function login(email: string, password: string): Promise<void> {
    if (!emailPattern.test(email)) {
      throw new Error("некорректный email"); // TODO
    }
    await AuthService.loginByEmail(email, password);
    setIsAuth(true);
  }

  async function register(data: RegisterData): Promise<void> {
    // validate
    if (!emailPattern.test(data.email)) {
      throw new Error("некорректный email"); // TODO
    }
    await AuthService.register(data);
    setIsAuth(true);
  }

  useEffect(() => {
    // setIsAuth(AuthService.isAuth());
  }, []);

  return {
    isAuth: isAuth,
    userId: AuthService.getUserId(),
    nickname: AuthService.getNickname(),
    login,
    register,
    logout,
  } as AuthData;
};

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const auth = useProvideAuth();
  return <authContext.Provider value={auth} children={children} />;
}
