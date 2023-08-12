package profile

import (
	"context"
	"time"
)

type UserModel struct {
	Id             int64     `json:"-"`
	Username       string    `json:"username"`
	CreatedAt      time.Time `json:"createdAt"`
	ActiveSessions int       `json:"activeSessions"`
	SessionLimit   int       `json:"sessionLimit"`
}

type FindOutput = UserModel

type Find struct {
	context.Context
}
