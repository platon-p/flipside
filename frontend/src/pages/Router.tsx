import { BrowserRouter, Route, Routes } from "react-router-dom";

import Main from "./MainPage/MainPage";
import EditSetPage from "./EditSetPage/EditSetPage";

import { Register } from "./Register";
import { CreateSetPage } from "./CreateSetPage/CreateSetPage";
import { LoginPage } from "./LoginPage/Login";
import { ViewSetPage } from "./ViewSetPage/CardSet";
import { TrainingPage } from "./TrainingPage/TrainingPage";

export function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/set/:slug" element={<ViewSetPage />} />
        <Route path="/set/:slug/edit" element={<EditSetPage />} />
        <Route path="/create-set" element={<CreateSetPage />} />
        <Route path="/training/:id" element={<TrainingPage />} />
      </Routes>
    </BrowserRouter>
  );
}
