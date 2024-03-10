import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Main, Register, Login, ViewSetPage, CreateSetPage, EditSetPage } from './components/pages';
import { AuthProvider } from './provider/AuthProvider';
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
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  )
}

export default App
