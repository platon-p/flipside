import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Main } from './components/pages/MainPage/MainPage'
import { Login } from './components/pages/LoginPage/Login'
import { Register } from './components/pages/Register'
import { ViewSetPage } from './components/pages/ViewSetPage/CardSet'
import { AuthProvider } from './provider/AuthProvider'
import { CreateSetPage } from './components/pages/CreateSetPage/CreateSetPage'

function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <Routes>
          <Route path='/' element={<Main />} />
          <Route path='/register' element={<Register />} />
          <Route path='/login' element={<Login />} />
          <Route path='/set/:slug' element={<ViewSetPage />} />
          <Route path='/create-set' element={<CreateSetPage />} />
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  )
}

export default App
