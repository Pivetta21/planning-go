package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/data/entity"
	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/google/uuid"
)

func (f *Refresh) Execute(opaqueToken uuid.UUID) (*RefreshOutput, error) {
	userSession, err := f.getByOpaqueToken(f.Context, opaqueToken)
	if err != nil {
		return nil, err
	}

	userSession.ExpiresAt = time.Now().UTC().Add(core.CookieDurationAuthSession)
	userSession.OpaqueToken = uuid.New()

	if err := f.refresh(f.Context, userSession); err != nil {
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

func (f *Refresh) getByOpaqueToken(ctx context.Context, opaqueToken uuid.UUID) (*entity.UserSession, error) {
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

	us.ExpiresAt = time.Now().UTC().Add(core.CookieDurationAuthSession)
	us.OpaqueToken = uuid.New()

	return &us, err
}

func (f *Refresh) refresh(ctx context.Context, e *entity.UserSession) error {
	queryCtx, cancel := context.WithTimeout(ctx, db.Ctx.DefaultTimeout)
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
