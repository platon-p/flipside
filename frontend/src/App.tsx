import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Main } from './components/pages/MainPage/MainPage'
import { Login } from './components/pages/Login'
import { Register } from './components/pages/Register'
import { CardSet } from './components/pages/CardSet'
import { AuthProvider } from './provider/AuthProvider'

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
