import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Main, Register, Login, ViewSetPage, CreateSetPage, EditSetPage, TrainingPage } from './pages';
import { AuthProvider } from './provider/AuthProvider';
import 'typeface-inter';
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <Routes>
          <Route path='/' element={<Main />} />
          <Route path='/register' element={<Register />} />
          <Route path='/login' element={<Login />} />
          <Route path='/set/:slug' element={<ViewSetPage />} />
          <Route path='/set/:slug/edit' element={<EditSetPage />} />
          <Route path='/create-set' element={<CreateSetPage />} />
          <Route path='/training/:id' element={<TrainingPage />} />
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  )
}

export default App
