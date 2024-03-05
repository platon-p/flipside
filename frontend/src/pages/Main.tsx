import { useNavigate } from "react-router-dom"
import { CardSetItem } from "../components/CardSetItem";
import { CardSetRepository } from "../repository/CardSetRepository";
import { useAuth } from "../service/AuthService";

export function Main() {
    const { isAuth } = useAuth();
    if (isAuth) {
        return <UserPage />
    }
    return <NewbiePage />
}

function UserPage() {
    const { userId, logout } = useAuth();
    const navigate = useNavigate();
    const cards = CardSetRepository.getCardSetsByOwner(userId!);

    function navigateToCardSet(slug: string) {
        navigate(`/set/${slug}`)
    }
    return <>
        <p>Hello</p>
        <p onClick={() => { logout() }}>logout</p>
        <div style={{
            display: 'flex',
            flexDirection: 'column',
            gap: '1em'
        }}>
            {cards.map((v, i) => {
                return <div key={i} onClick={() => navigateToCardSet(v.slug)}>
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