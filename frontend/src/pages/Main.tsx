import { useNavigate } from "react-router-dom"
import { CardSetItem } from "../components/CardSetItem";

export function Main() {
    const auth = true;
    return <>
        {auth ? <UserPage /> : <NewbiePage />}
    </>
}

function UserPage() {
    const navigate = useNavigate();
    const cards = [
        { title: "1. CardSet title", slug: "set1" },
        { title: "2. CardSet title", slug: "set2" }
    ]
    function navigateToCardSet(slug: string) {
        navigate(`/set/${slug}`)
    }
    return <>
        <p>Hello</p>
        <div style={{
            display: 'flex',
            flexDirection: 'column',
            gap: '1em'
        }}>
            {cards.map((v, i) => {
                return <div onClick={_ => navigateToCardSet(v.slug)}>
                    <CardSetItem title={v.title} slug={v.slug} key={i} />
                </div>
            })}
        </div>
    </>
}

function NewbiePage() {
    const navigate = useNavigate()

    function goToLogin() {
        navigate('/login')
    }
    function goToRegister() {
        navigate('/register')
    }


    return <>
        <h1>Index</h1 >
        <p onClick={goToLogin}>Вход</p>
        <p onClick={goToRegister}>Регистрация</p>
    </>
}