import { useNavigate } from "react-router-dom"
import { Button } from "@/shared/Button"
import { Input } from "@/shared/Input"
import { useAuth } from "@/hooks/Auth"
import { useState } from "react";
import './Login.css';

export function LoginPage() {
  const [view, setView] = useState<'login' | 'register'>('login');

  return <div>
    <div className='container' style={{ height: 700, left: 0 }}>
      <ButtonBox view={view} onChange={setView} />
      <Login active={view === 'login'} />
      <Register active={view === 'register'} />
    </div>
  </div>
}

function ButtonBox({view, onChange}: {view: 'login' | 'register', onChange: (view: 'login' | 'register') => void}) {
  return <div className='button-box'>
    <div className="underline"></div>
    <div style={{
      left: view === 'login' ? 0 : '43%',
      width: view === 'login' ? '43%' : '57%',
      backgroundColor: '#F1694F'
    }}></div>
    <button className="toggle-btn log"
      onClick={() => onChange('login')}
      style={{
        color: view === 'login' ? '#F1694F' : undefined
      }}>вход</button>
    <button className="toggle-btn reg" onClick={() => onChange('register')}
      style={{
        color: view === 'register' ? '#F1694F' : undefined
      }}>регистрация</button>
  </div>
}

function Register({ active }: { active: boolean }) {
  const navigate = useNavigate();
  const { register } = useAuth();
  const [formValues, setFormValues] = useState({
    name: '',
    nickname: '',
    email: '',
    password: '',
    repeatPassword: ''
  });
  const [errorRegMessage, setErrorRegMessage] = useState("")

  function handleInput(e: React.FormEvent<HTMLInputElement>, key: keyof typeof formValues) {
    setFormValues({
      ...formValues,
      [key]: e.currentTarget.value
    });
  }

  function goToMain() {
    navigate('/')
  }

  async function submitreg() {
    // validate
    if (formValues.password !== formValues.repeatPassword) {
      setErrorRegMessage('пароли не совпадают');
      return;
    }
    // register
    const res = await register(formValues);
    if (res) {
      setErrorRegMessage(res);
      return;
    }
    goToMain();
  }

  return <div className='input-group form' onSubmit={(e) => { e.preventDefault() }} style={{
    left: active ? 0 : 450
  }}>
    <div>
      <Input
        placeholder="имя"
        onInput={(e) => handleInput(e, 'name')}
        className='authorize' />
    </div>
    <div>
      <Input
        placeholder="никнейм"
        onInput={(e) => handleInput(e, 'nickname')}
        className='authorize' />

    </div>
    <div>
      <Input placeholder="почта"
        onInput={(e) => handleInput(e, 'email')}
        className='authorize' />

    </div>
    <div>
      <Input placeholder="пароль"
        type="password" onInput={(e) => handleInput(e, 'password')}
        className='authorize' />

    </div>
    <div>
      <Input
        placeholder="повторите пароль"
        type="password" onInput={(e) => handleInput(e, 'repeatPassword')}
        className='authorize' />

    </div>
    {errorRegMessage && <p style={{ color: 'red', fontFamily: 'inter-norm' }}>{errorRegMessage}</p>}
    <Button className="create-set regbtn" onClick={submitreg}>зарегистрироваться</Button>


  </div>
}

function Login({ active }: { active: boolean }) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorLoginMessage, setErrorLoginMessage] = useState("")
  const auth = useAuth();
  const navigate = useNavigate();

  function goToMain() {
    navigate('/')
  }

  async function submit() {
    const res = await auth.login(email, password);
    if (res) {
      setErrorLoginMessage(res);
      return
    }
    goToMain();
  }

  return <div className='input-group form' style={{
    left: active ? 0 : '-400px'
  }}>
    <div>
      <Input className='authorize' onInput={(e) => setEmail(e.currentTarget.value)} placeholder="почта" />
    </div>
    <div>
      <Input className='authorize' onInput={(e) => setPassword(e.currentTarget.value)} placeholder="пароль" type="password" />
    </div>
    {errorLoginMessage && <p style={{ color: 'red', fontFamily: 'inter-norm' }}>{errorLoginMessage}</p>}
    <Button className="create-set logbtn" onClick={submit}>войти</Button>

  </div>
}

export default LoginPage