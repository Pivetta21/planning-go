package session

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/infra/db"
)

func (f *SessionDelete) Execute(identifier string) error {
	loggedUser := core.GetLoggedUser(f.Context)

	if loggedUser.Session.Identifier == identifier {
		return errors.New("cannot delete current session, please logout")
	}

	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	var deletedId int64
	row := db.Ctx.Conn.QueryRowContext(
		queryCtx,
		`
		DELETE FROM public.user_sessions 
		WHERE user_id = $1 AND identifier = $2
		RETURNING id
		`,
		loggedUser.Id, identifier,
	)

	if err := row.Scan(&deletedId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("the requested session could not be deleted")
		}

		return err
	}

	return nil
}
