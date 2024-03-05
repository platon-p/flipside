import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Main } from './pages/Main'
import { Login } from './pages/Login'
import { Register } from './pages/Register'
import { CardSet } from './pages/CardSet'
import { AuthProvider } from './service/AuthService'

function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <Routes>
          <Route path='/' element={<Main />} />
          <Route path='/register' element={<Register />} />
          <Route path='/login' element={<Login />} />
          <Route path='/set/:slug' element={<CardSet />} />
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  )
}

export default App
