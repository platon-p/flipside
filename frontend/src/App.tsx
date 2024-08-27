import { AuthProvider } from "./hooks/AuthProvider";
import { Router } from "./pages/Router";
import "typeface-inter";
import "./App.css";

function App() {
  return (
    <AuthProvider>
      <Router />
    </AuthProvider>
  );
}

export default App;
