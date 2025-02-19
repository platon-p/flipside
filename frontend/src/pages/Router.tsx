import { BrowserRouter, Route, Routes } from "react-router-dom";

import {
  CreateSet,
  Login,
  ViewSet,
  Training,
  Main,
  EditSet,
  Profile,
} from "./index";

export function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/login" element={<Login />} />
        <Route path="/set/:slug" element={<ViewSet />} />
        <Route path="/set/:slug/edit" element={<EditSet />} />
        <Route path="/create-set" element={<CreateSet />} />
        <Route path="/training/:id" element={<Training />} />
        <Route path="/profile/:nickname" element={<Profile />} />
      </Routes>
    </BrowserRouter>
  );
}
