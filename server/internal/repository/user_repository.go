package repository

import (
	"context"
	"strings"

	"github.com/Pivetta21/planning-go/internal/data/entity"
	"github.com/Pivetta21/planning-go/internal/infra/db"
)

type IUserRepository interface {
	Save(c context.Context, e *entity.User) (int64, error)
	GetByUsername(c context.Context, username string) (*entity.User, error)
}

type UserRepository struct {
	Db *db.DbContext
}

func NewUserRepository(db *db.DbContext) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, e *entity.User) (int64, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.Db.DefaultTimeout)
	defer cancel()

	row := r.Db.Conn.QueryRowContext(
		queryCtx,
		`
		INSERT INTO public.users (username, password)
		VALUES ($1, $2) 
		RETURNING id
		`,
		e.Username, e.Password,
	)

	var lastInsertedId int64
	if err := row.Scan(&lastInsertedId); err != nil {
		return 0, err
	}

	return lastInsertedId, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.Db.DefaultTimeout)
	defer cancel()

	row := r.Db.Conn.QueryRowContext(
		queryCtx,
		`
		SELECT id, username, password, created_at 
		FROM public.users 
		WHERE username = $1
		`,
		strings.ToLower(username),
	)

	var user entity.User
	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
