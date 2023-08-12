package core

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type LoggedUser struct {
	Id           int64
	Username     string
	CreatedAt    time.Time
	SessionLimit int
	Session      LoggedUserSession
}

type LoggedUserSession struct {
	Id          int64
	Identifier  string
	OpaqueToken uuid.UUID
	ExpiresAt   time.Time
}

func GetLoggedUser(ctx context.Context) *LoggedUser {
	return ctx.Value(ContextKeyLoggedUser).(*LoggedUser)
}

func (c *LoggedUser) IsSessionExpired() bool {
	return c.Session.ExpiresAt.Before(time.Now().UTC())
}
