package auth

import (
	"net/http"
	"time"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/google/uuid"
)

func (u *Refresh) Execute(opaqueToken uuid.UUID) (*RefreshOutput, error) {
	userSession, err := u.UserSessionRepository.GetByOpaqueToken(u.Context, opaqueToken)
	if err != nil {
		return nil, err
	}

	userSession.ExpiresAt = time.Now().UTC().Add(core.CookieDurationAuthSession)
	userSession.OpaqueToken = uuid.New()

	if err := u.UserSessionRepository.Refresh(u.Context, userSession); err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     core.CookieNameAuthSession.String(),
		Value:    userSession.OpaqueToken.String(),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   int(core.CookieDurationAuthSession.Seconds()),
	}

	out := &RefreshOutput{
		Cookie: cookie,
	}

	return out, nil
}
