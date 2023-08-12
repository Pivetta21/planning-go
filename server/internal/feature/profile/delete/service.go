package update

import (
	"context"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/infra/db"
)

func (f Delete) Execute() error {
	loggedUser := core.GetLoggedUser(f.Context)

	if err := f.deleteUserById(loggedUser.Id); err != nil {
		return err
	}

	return nil
}

func (f Delete) deleteUserById(userId int64) error {
	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	_, err := db.Ctx.Conn.ExecContext(
		queryCtx,
		`
		DELETE FROM public.users
		WHERE id = $1
		`,
		userId,
	)

	return err
}
