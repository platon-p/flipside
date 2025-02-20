package service

import (
	"github.com/platon-p/flipside/authservice/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockRefreshTokenRepository struct {
	CreateFn      func(userId int, token string, expiresAt time.Time) (*model.RefreshToken, error)
	FindByTokenFn func(token string) (*model.RefreshToken, error)
	FindByUserFn  func(userId int) (*model.RefreshToken, error)
	DeleteFn      func(token string) error
}

func (m *mockRefreshTokenRepository) Create(userId int, token string, expiresAt time.Time) (*model.RefreshToken, error) {
	return m.CreateFn(userId, token, expiresAt)
}

func (m *mockRefreshTokenRepository) FindByToken(token string) (*model.RefreshToken, error) {
	return m.FindByTokenFn(token)
}

func (m *mockRefreshTokenRepository) FindByUser(userId int) (*model.RefreshToken, error) {
	return m.FindByUserFn(userId)
}

func (m *mockRefreshTokenRepository) Delete(token string) error {
	return m.DeleteFn(token)
}

func TestRefreshTokenService_CreateToken_NonExistingToken(t *testing.T) {
	var newToken model.RefreshToken
	repo := mockRefreshTokenRepository{
		CreateFn: func(userId int, token string, expiresAt time.Time) (*model.RefreshToken, error) {
			newToken = model.RefreshToken{User: &model.User{Id: userId}, Token: token, ExpiresAt: expiresAt}
			return &newToken, nil
		},
		FindByUserFn: func(userId int) (*model.RefreshToken, error) {
			return nil, nil
		},
	}
	serv := NewRefreshTokenService(&repo, 3*time.Second)
	user := model.User{Id: 1}
	res, err := serv.CreateToken(&user)
	assert.NoError(t, err)
	assert.Equal(t, &newToken, res)
}

func TestRefreshTokenService_CreateToken_ExistingToken(t *testing.T) {
	deleteCalled := false
	user := model.User{Id: 1}
	oldToken := model.RefreshToken{User: &user, Token: "old-token", ExpiresAt: time.Now().Add(-time.Second)}
	repo := &mockRefreshTokenRepository{
		DeleteFn: func(token string) error {
			deleteCalled = true
			return nil
		},
		CreateFn: func(userId int, token string, expiresAt time.Time) (*model.RefreshToken, error) {
			return &model.RefreshToken{User: &user, Token: token, ExpiresAt: expiresAt}, nil
		},
		FindByUserFn: func(userId int) (*model.RefreshToken, error) {
			return &oldToken, nil
		},
	}
	s := NewRefreshTokenService(repo, 3*time.Second)
	token, err := s.CreateToken(&user)
	assert.True(t, deleteCalled)
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, &user, token.User)
	assert.InDelta(t, time.Now().Add(3*time.Second).UnixMilli(), token.ExpiresAt.UnixMilli(), float64(time.Millisecond))
}

func TestRefreshTokenService_CheckToken(t *testing.T) {
	user := model.User{Id: 1}
	existingToken := model.RefreshToken{Token: "token", ExpiresAt: time.Now().Add(300 * time.Millisecond), User: &user}
	repo := &mockRefreshTokenRepository{
		FindByTokenFn: func(token string) (*model.RefreshToken, error) {
			if token == existingToken.Token {
				return &existingToken, nil
			}
			return nil, nil
		},
	}
	serv := NewRefreshTokenService(repo, 3*time.Second)
	res, err := serv.CheckToken(existingToken.Token)
	assert.NoError(t, err)
	assert.Equal(t, &user, res)

	// non existing token
	res, err = serv.CheckToken("non-existing-token")
	assert.ErrorIs(t, err, InvalidRefreshToken)

	// expired token
	time.Sleep(500 * time.Millisecond)
	res, err = serv.CheckToken(existingToken.Token)
	assert.ErrorIs(t, err, ExpiredRefreshToken)
}
