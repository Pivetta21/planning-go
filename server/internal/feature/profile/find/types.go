package profile

import (
	"context"
	"time"
)

type UserModel struct {
	Username       string    `json:"username"`
	CreatedAt      time.Time `json:"createdAt"`
	SessionLimit   int       `json:"sessionLimit"`
	ActiveSessions int       `json:"activeSessions"`
}

type FindOutput = UserModel

type Find struct {
	context.Context
}
