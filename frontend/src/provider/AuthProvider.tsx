import { useState } from "react";
import { AuthData, authContext } from "../hooks/Auth";

const useProvideAuth = (): AuthData => {
    const [isAuth, setIsAuth] = useState(false);
    const [userId] = useState<number | null>(null);

    function logout() {
        setIsAuth(true);
    }

    const init = () => {

    }

    return {
        isAuth,
        userId,
        logout,
    }
}

export function AuthProvider({ children }: { children: React.ReactNode }) {
    const auth = useProvideAuth();
    return <authContext.Provider value={auth} children={children} />
}
