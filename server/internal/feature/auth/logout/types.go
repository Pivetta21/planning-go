package auth

import (
	"context"
	"net/http"

	"github.com/Pivetta21/planning-go/internal/repository"
)

type LogoutOutput struct {
	Cookie *http.Cookie `json:"-"`
}

type Logout struct {
	Context               context.Context
	UserSessionRepository repository.IUserSessionRepository
}

func NewLogout(ctx context.Context, userSessionRepository *repository.UserSessionRepository) *Logout {
	return &Logout{
		Context:               ctx,
		UserSessionRepository: userSessionRepository,
	}
}
