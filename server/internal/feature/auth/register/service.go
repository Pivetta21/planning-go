package auth

import (
	"context"
	"fmt"

	"github.com/Pivetta21/planning-go/internal/data/entity"
	"github.com/Pivetta21/planning-go/internal/infra/db"
	"golang.org/x/crypto/bcrypt"
)

func (f *Register) Execute(in *RegisterInput) (*RegisterOutput, error) {
	hashedPassword, err := f.hashPassword(in.Password)
	if err != nil {
		return nil, err
	}

	user, err := entity.NewUser(0, in.Username, hashedPassword)
	if err != nil {
		return nil, err
	}

	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	row := db.Ctx.Conn.QueryRowContext(
		queryCtx,
		`
		INSERT INTO public.users (username, password)
		VALUES ($1, $2) 
		RETURNING id
		`,
		user.Username, user.Password,
	)

	var lastInsertedId int64
	if err := row.Scan(&lastInsertedId); err != nil {
		return nil, err
	}

	out := &RegisterOutput{
		Message: fmt.Sprintf("User %#q created successfully", in.Username),
	}

	return out, err
}

func (f *Register) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}
