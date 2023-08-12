package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/Pivetta21/planning-go/internal/data/enum"
)

type UserSessionDto struct {
	Id        int64
	Active    bool
	CreatedAt time.Time
}

type Input struct {
	Username string             `json:"username"`
	Password string             `json:"password"`
	Origin   enum.SessionOrigin `json:"origin"`
}

type Output struct {
	Message string       `json:"message"`
	Cookie  *http.Cookie `json:"-"`
}

type Login struct {
	context.Context
}
