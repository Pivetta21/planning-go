package session

import (
	"context"
	"time"

	"github.com/Pivetta21/planning-go/internal/data/enum"
)

type UserSessionModel struct {
	Id                int64              `json:"-"`
	Identifier        string             `json:"identifier"`
	Origin            enum.SessionOrigin `json:"-"`
	OriginDescription string             `json:"origin"`
	CreatedAt         time.Time          `json:"createdAt"`
	Active            bool               `json:"active"`
	Current           bool               `json:"current"`
}

type Output []UserSessionModel

type List struct {
	context.Context
}
