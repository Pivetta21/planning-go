package repository

import (
	"context"

	"github.com/Pivetta21/planning-go/internal/data/entity"
	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/google/uuid"
)

type IUserSessionRepository interface {
	Save(ctx context.Context, e *entity.UserSession) (int64, error)
	GetByOpaqueToken(ctx context.Context, opaqueToken uuid.UUID) (*entity.UserSession, error)
	DeleteByOpaqueToken(ctx context.Context, opaqueToken uuid.UUID) error
	Refresh(ctx context.Context, e *entity.UserSession) error
}

type UserSessionRepository struct {
	DB *db.DbContext
}

func NewUserSessionRepository(db *db.DbContext) *UserSessionRepository {
	return &UserSessionRepository{
		DB: db,
	}
}

func (r *UserSessionRepository) Save(ctx context.Context, e *entity.UserSession) (int64, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.DB.DefaultTimeout)
	defer cancel()

	row := r.DB.Conn.QueryRowContext(
		queryCtx,
		`
		INSERT INTO public.user_sessions(user_id, identifier, opaque_token, origin, expires_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
		`,
		e.UserId, e.Identifier, e.OpaqueToken, e.Origin, e.ExpiresAt,
	)

	var lastInsertedId int64
	if err := row.Scan(&lastInsertedId); err != nil {
		return 0, err
	}

	return lastInsertedId, nil
}

func (r *UserSessionRepository) GetByOpaqueToken(ctx context.Context, opaqueToken uuid.UUID) (*entity.UserSession, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.DB.DefaultTimeout)
	defer cancel()

	row := r.DB.Conn.QueryRowContext(
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

	return &us, nil
}

func (r *UserSessionRepository) DeleteByOpaqueToken(ctx context.Context, opaqueToken uuid.UUID) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.DB.DefaultTimeout)
	defer cancel()

	_, err := r.DB.Conn.ExecContext(
		queryCtx,
		`DELETE FROM public.user_sessions WHERE opaque_token = $1`,
		opaqueToken,
	)

	return err
}

func (r *UserSessionRepository) Refresh(ctx context.Context, e *entity.UserSession) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.DB.DefaultTimeout)
	defer cancel()

	_, err := r.DB.Conn.ExecContext(
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
