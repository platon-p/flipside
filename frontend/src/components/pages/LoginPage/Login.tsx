import { useNavigate } from "react-router-dom"
import { Button } from "@/components/shared/Button"
import { Input } from "@/components/shared/Input"
import { useAuth } from "@/hooks/Auth"
import { useState } from "react";
import './Login.css';

export function Login() {
  const { isAuth, register, login } = useAuth();
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorLoginMessage, setErrorLoginMessage] = useState("")
  const [errorRegMessage, setErrorRegMessage] = useState("")

  const [formValues, setFormValues] = useState({
    name: '',
    nickname: '',
    email: '',
    password: '',
    repeatPassword: ''
  });

  function goToLogin() {
    navigate('/login')
  }

  function goToMain() {
    navigate('/')
  }

  function handleInput(e: React.FormEvent<HTMLInputElement>, key: keyof typeof formValues) {
    setFormValues({
      ...formValues,
      [key]: e.currentTarget.value
    });
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

  if (isAuth) {
    goToMain();
  }

  function goToRegister() {
    navigate('/register');
  }

  async function submit() {
    const res = await login(email, password);
    if (res) {
      setErrorLoginMessage(res);
      return
    }
    goToMain();
  }

  if (isAuth) {
    goToMain();
  }

  var loginn = document.getElementById('login')!;
  var registers = document.getElementById('register')!;
  var btn = document.getElementById('btn')!;
  var b_r = document.getElementById('b_r')!;
  var b_l = document.getElementById('b_l')!;
  const logined = () => {
    login.style.left = "0";
    registers.style.left = "450px";
    btn.style.left = "0";
    btn.style.width = "34vw";
    b_l.style.color = "white";
    b_r.style.color = "black";
  };

  const registered = () => {
    login.style.left = "-400px";
    registers.style.left = "0";
    btn.style.left = "27.57vw";
    btn.style.width = "48.97vw";
    b_r.style.color = "white";
    b_l.style.color = "black";
  };

  return <>
    <div className='container' style={{ height: 700, left: 0 }}>
      <style>{'body{ align-items: flex-start;}'}</style>
      <div className='button-box'>
        <div id="btn"></div>
        <button id='b_l' className="toggle-btn log"
          onClick={() => logined()}
          style={{ color: 'white' }}>вход</button>
        <button id='b_r'
          className="toggle-btn reg" onClick={() => registered()}
          style={{ color: 'black' }}>регистрация</button>
      </div>
      <form onSubmit={(e) => { e.preventDefault() }}
        id='login' className='input-group' style={{ left: 0 }}>
        <div>
          <Input className='authorize' onInput={(e: any) => setEmail(e.currentTarget.value)} placeholder="почта" />
        </div>
        <div>
          <Input className='authorize' onInput={(e: any) => setPassword(e.currentTarget.value)} placeholder="пароль" type="password" />
        </div>
        {errorLoginMessage && <p style={{ color: 'red' }}>{errorLoginMessage?.toString()}</p>}
        <Button onClick={submit} className="enter" style={{ borderRadius: 28, color: "white" }}>
          <style>{'button { background-color: rgba(241, 105, 79, 1); }'}</style>
          ВОЙТИ
        </Button>
      </form>
      <form
        id='register' className='input-group'
        onSubmit={(e) => { e.preventDefault() }}>
        <div>
          <Input
            placeholder="имя"
            onInput={(e: any) => handleInput(e, 'name')}
            className='authorize' />
        </div>
        <div>
          <Input
            placeholder="никнейм"
            onInput={(e: any) => handleInput(e, 'nickname')}
            className='authorize' />

        </div>
        <div>
          <Input placeholder="почта"
            onInput={(e: any) => handleInput(e, 'email')}
            className='authorize' />

        </div>
        <div>
          <Input placeholder="пароль"
            type="password" onInput={(e: any) => handleInput(e, 'password')}
            className='authorize' />

        </div>
        <div>
          <Input
            placeholder="повторите пароль"
            type="password" onInput={(e: any) => handleInput(e, 'repeatPassword')}
            className='authorize' />

        </div>
        {errorRegMessage && <p style={{ color: 'red', fontFamily: 'inter-norm' }}>{errorRegMessage}</p>}
        <Button onClick={submitreg}
          className='enter'
          type="submit"
          style={{ borderRadius: 28, color: "white" }}>
          <style>{'button { background-color: rgba(241, 105, 79, 1); }'}</style>
          ЗАРЕГИСТРИРОВАТЬСЯ
        </Button>

      </form>
    </div>
  </>
}

export default Login