package session

import (
	"context"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/infra/db"
)

func (f *List) Execute() (Output, error) {
	loggedUser := core.GetLoggedUser(f.Context)

	userSessions, err := f.listByUserId(loggedUser)
	if err != nil {
		return nil, err
	}

	return userSessions, nil
}

func (f *List) listByUserId(loggedUser *core.LoggedUser) ([]UserSessionModel, error) {
	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	rows, err := db.Ctx.Conn.QueryContext(
		queryCtx,
		`
		SELECT id, identifier, origin, now() <= expires_at AS active, created_at, CASE WHEN id = $2 THEN true ELSE false END AS current 
		FROM public.user_sessions
		WHERE user_id = $1
		`,
		loggedUser.Id, loggedUser.Session.Id,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userSessions []UserSessionModel
	for rows.Next() {
		var us UserSessionModel
		if err := rows.Scan(
			&us.Id,
			&us.Identifier,
			&us.Origin,
			&us.Active,
			&us.CreatedAt,
			&us.Current,
		); err != nil {
			return userSessions, err
		}

		us.OriginDescription = us.Origin.String()

		userSessions = append(userSessions, us)
	}

	if err := rows.Err(); err != nil {
		return userSessions, err
	}

	return userSessions, nil
}
