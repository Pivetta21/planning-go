package auth

import (
	"context"
	"net/http"

	"github.com/Pivetta21/planning-go/internal/data/enum"
	"github.com/Pivetta21/planning-go/internal/repository"
)

type LoginInput struct {
	Username string             `json:"username"`
	Password string             `json:"password"`
	Origin   enum.SessionOrigin `json:"origin"`
}

type LoginOutput struct {
	Message string       `json:"message"`
	Cookie  *http.Cookie `json:"-"`
}

type Login struct {
	Context               context.Context
	UserRepository        repository.IUserRepository
	UserSessionRepository repository.IUserSessionRepository
}

func NewLogin(
	ctx context.Context,
	userRepository *repository.UserRepository,
	userSessionRepository *repository.UserSessionRepository,
) *Login {
	return &Login{
		Context:               ctx,
		UserRepository:        userRepository,
		UserSessionRepository: userSessionRepository,
	}
}
