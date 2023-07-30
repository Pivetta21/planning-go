package session

import (
	"time"

	"github.com/Pivetta21/planning-go/internal/data/enum"
)

type SessionModel struct {
	Id                int64              `json:"-"`
	Identifier        string             `json:"identifier"`
	ExpiresAt         time.Time          `json:"expiresAt"`
	Origin            enum.SessionOrigin `json:"-"`
	OriginDescription string             `json:"origin"`
	CreatedAt         time.Time          `json:"createdAt"`
}

type SessionListOutput []SessionModel
