package transport

import (
	"context"
	"net/http"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/google/uuid"
)

func AuthMiddleware(next http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		authSessionCookie, err := r.Cookie(core.CookieNameAuthSession.String())
		if err != nil || authSessionCookie.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		opaqueToken, err := uuid.Parse(authSessionCookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		loggedUser, err := fetchLoggedUser(r.Context(), opaqueToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if loggedUser.IsSessionExpired() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), core.ContextKeyLoggedUser, loggedUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(middleware)
}

func fetchLoggedUser(ctx context.Context, opaqueToken uuid.UUID) (*core.LoggedUser, error) {
	queryCtx, cancel := context.WithTimeout(ctx, db.Ctx.DefaultTimeout)
	defer cancel()

	row := db.Ctx.Conn.QueryRowContext(
		queryCtx,
		`
		SELECT 
		    u.id, u.username, u.created_at, u.session_limit,
		    us.id, us.identifier, us.opaque_token, us.expires_at
		FROM public.user_sessions AS us
		JOIN public.users AS u ON u.id = us.user_id
		WHERE us.opaque_token = $1
		`,
		opaqueToken,
	)

	var user core.LoggedUser
	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.CreatedAt,
		&user.SessionLimit,
		&user.Session.Id,
		&user.Session.Identifier,
		&user.Session.OpaqueToken,
		&user.Session.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
