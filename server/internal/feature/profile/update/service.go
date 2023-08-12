package update

import (
	"context"
	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/infra/db"
)

func (f Update) Execute(in Input) error {
	loggedUser := core.GetLoggedUser(f.Context)
	return f.updateLoggerUser(loggedUser.Id, in)
}

func (f Update) updateLoggerUser(userId int64, in Input) error {
	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	_, err := db.Ctx.Conn.ExecContext(
		queryCtx,
		`
		UPDATE public.users 
		SET username = $2
		WHERE id = $1
		`,
		userId, in.Username,
	)

	return err
}
