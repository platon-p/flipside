import { useEffect, useState } from "react";
import { AuthData, RegisterData, authContext } from "../hooks/Auth";
import { AuthService } from "../service/AuthService";

const useProvideAuth = (): AuthData => {
    const [isAuth, setIsAuth] = useState(false);
    const [userId] = useState<number | null>(null);

    function logout() {
        AuthService.logout();
        setIsAuth(false);
    }

    async function login(email: string, password: string): Promise<string | null> {
        const res = await AuthService.loginByEmail(email, password);
        if (!res) {
            setIsAuth(true);
        }
        return res;
    }

    async function register(data: RegisterData): Promise<string | null> {
        // validate
        const emailPattern = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]$/;
        if (!emailPattern.test(data.email)) {
            return 'Некорректный email';
        }
        const res = await AuthService.register(data);
        if (!res) {
            setIsAuth(true);
        }
        return res;
    }

    const init = () => {
        console.log('init');
        setIsAuth(AuthService.isAuth());
    }

    useEffect(() => {
        init();
    }, []);

    return {
        isAuth,
        userId,
        login,
        register,
        logout,
    }
}

export function AuthProvider({ children }: { children: React.ReactNode }) {
    const auth = useProvideAuth();
    return <authContext.Provider value={auth} children={children} />
}
