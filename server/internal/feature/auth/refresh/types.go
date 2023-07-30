package auth

import (
	"context"
	"net/http"

	"github.com/Pivetta21/planning-go/internal/repository"
)

type RefreshOutput struct {
	Cookie *http.Cookie `json:"-"`
}

type Refresh struct {
	Context               context.Context
	UserSessionRepository repository.IUserSessionRepository
}

func NewRefresh(ctx context.Context, userSessionRepository *repository.UserSessionRepository) *Refresh {
	return &Refresh{
		Context:               ctx,
		UserSessionRepository: userSessionRepository,
	}
}
