package profile

import (
	"context"
	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/infra/db"
)

func (f *Find) Execute() (*FindOutput, error) {
	user, err := f.populateUser()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (f *Find) populateUser() (*UserModel, error) {
	loggedUser := core.GetLoggedUser(f.Context)

	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	row := db.Ctx.Conn.QueryRowContext(
		queryCtx,
		`
		SELECT count(id)
		FROM public.user_sessions
		WHERE user_id = $1
		`,
		loggedUser.Id,
	)

	var activeSessions int
	if err := row.Scan(&activeSessions); err != nil {
		return nil, err
	}

	user := &UserModel{
		Username:       loggedUser.Username,
		CreatedAt:      loggedUser.CreatedAt,
		SessionLimit:   loggedUser.SessionLimit,
		ActiveSessions: activeSessions,
	}

	return user, nil
}
