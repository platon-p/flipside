package service

type AuthService struct {
}

type TokenPair struct {
	AccessToken  string
	TokenType    string
	RefreshToken string
}

func (s *AuthService) Register(name, nickname, email, password string) (*TokenPair, error) 

func (s *AuthService) LoginByEmail(email, password string) (*TokenPair, error)

func (s *AuthService) LoginByToken(refreshToken string) (*TokenPair, error)
