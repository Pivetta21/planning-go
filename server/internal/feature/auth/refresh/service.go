package auth

import (
	"context"
	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/data/entity"
	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/google/uuid"
	"net/http"
)

func (f *Refresh) Execute(opaqueToken uuid.UUID) (*Output, error) {
	userSession, err := f.getByOpaqueToken(opaqueToken)
	if err != nil {
		return nil, err
	}

	if err := f.refresh(userSession); err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     core.CookieNameAuthSession.String(),
		Value:    userSession.OpaqueToken.String(),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   core.TimeExpirationAuthCookie.MaxAge(),
	}

	out := &Output{
		Cookie: cookie,
	}

	return out, nil
}

func (f *Refresh) getByOpaqueToken(opaqueToken uuid.UUID) (*entity.UserSession, error) {
	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	row := db.Ctx.Conn.QueryRowContext(
		queryCtx,
		`
		SELECT id, user_id, identifier, opaque_token, origin, expires_at, created_at
		FROM public.user_sessions 
		WHERE opaque_token = $1
		`,
		opaqueToken,
	)

	var us entity.UserSession
	err := row.Scan(
		&us.Id,
		&us.UserId,
		&us.Identifier,
		&us.OpaqueToken,
		&us.Origin,
		&us.ExpiresAt,
		&us.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	us.ExpiresAt = core.TimeExpirationAuthSession.Future()
	us.OpaqueToken = uuid.New()

	return &us, err
}

func (f *Refresh) refresh(e *entity.UserSession) error {
	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	_, err := db.Ctx.Conn.ExecContext(
		queryCtx,
		`
		UPDATE public.user_sessions 
		SET expires_at = $2, opaque_token = $3
		WHERE id = $1
		`,
		e.Id, e.ExpiresAt, e.OpaqueToken,
	)

	return err
}
