import { useAuth } from "@/hooks/Auth";
import { NewbiePage } from "./NewbiePage/NewbiePage";
import { UserPage } from "./UserPage/UserPage";

export function Main() {
  const auth = useAuth();
  if (auth.isAuth) {
    return <UserPage />;
  }
  return <NewbiePage />;
}

export default Main;
