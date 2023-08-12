package auth

import (
	"context"
	"net/http"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/google/uuid"
)

func (f *Logout) Execute() (*Output, error) {
	loggedUser := core.GetLoggedUser(f.Context)

	err := f.deleteByOpaqueToken(loggedUser.Session.OpaqueToken)
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

	out := &Output{
		Cookie: cookie,
	}

	return out, nil
}

func (f *Logout) deleteByOpaqueToken(opaqueToken uuid.UUID) error {
	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	_, err := db.Ctx.Conn.ExecContext(
		queryCtx,
		`DELETE FROM public.user_sessions WHERE opaque_token = $1`,
		opaqueToken,
	)

	return err
}
