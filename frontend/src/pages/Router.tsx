import { BrowserRouter, Route, Routes } from "react-router-dom";

import {
  Register,
  CreateSetPage,
  Login,
  ViewSetPage,
  TrainingPage,
  Main,
  EditSetPage,
  Profile,
} from "./index";

export function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/set/:slug" element={<ViewSetPage />} />
        <Route path="/set/:slug/edit" element={<EditSetPage />} />
        <Route path="/create-set" element={<CreateSetPage />} />
        <Route path="/training/:id" element={<TrainingPage />} />
        <Route path="/profile/:nickname" element={<Profile />} />
      </Routes>
    </BrowserRouter>
  );
}
