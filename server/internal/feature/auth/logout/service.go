package auth

import (
	"net/http"

	"github.com/Pivetta21/planning-go/internal/core"
)

func (u *Logout) Execute() (*LogoutOutput, error) {
	loggedUser := core.GetLoggedUser(u.Context)

	err := u.UserSessionRepository.DeleteByOpaqueToken(u.Context, loggedUser.Session.OpaqueToken)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     core.CookieNameAuthSession.String(),
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   0,
	}

	out := &LogoutOutput{
		Cookie: cookie,
	}

	return out, nil
}
