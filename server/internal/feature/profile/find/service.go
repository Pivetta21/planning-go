package profile

import (
	"context"
	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/infra/db"
)

func (f *Find) Execute() (*FindOutput, error) {
	loggedUser := core.GetLoggedUser(f.Context)

	user, err := f.fetchUserById(loggedUser.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (f *Find) fetchUserById(userId int64) (*UserModel, error) {
	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	row := db.Ctx.Conn.QueryRowContext(
		queryCtx,
		`
		SELECT id, username, created_at, session_limit
		FROM public.users
		WHERE id = $1
		`,
		userId,
	)

	var user UserModel
	if err := row.Scan(
		&user.Id, &user.Username, &user.CreatedAt, &user.SessionLimit,
	); err != nil {
		return nil, err
	}

	return f.fetchUserActiveSessions(&user)
}

func (f *Find) fetchUserActiveSessions(user *UserModel) (*UserModel, error) {
	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	row := db.Ctx.Conn.QueryRowContext(
		queryCtx,
		`
		SELECT count(id)
		FROM public.user_sessions
		WHERE user_id = $1
		`,
		user.Id,
	)

	var activeSessions int
	if err := row.Scan(&activeSessions); err != nil {
		return nil, err
	}

	user.ActiveSessions = activeSessions
	return user, nil
}
