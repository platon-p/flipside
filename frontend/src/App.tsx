import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Main } from './pages/Main'
import { Login } from './pages/Login'
import { Register } from './pages/Register'
import { CardSet } from './pages/CardSet'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Main />} />
        <Route path='/register' element={<Register />} />
        <Route path='/login' element={<Login />} />
        <Route path='/set/:slug' element={<CardSet />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
