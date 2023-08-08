package auth

import (
	"context"
	"net/http"

	"github.com/Pivetta21/planning-go/internal/data/enum"
)

type LoginInput struct {
	Username string             `json:"username"`
	Password string             `json:"password"`
	Origin   enum.SessionOrigin `json:"origin"`
}

type LoginOutput struct {
	Message string       `json:"message"`
	Cookie  *http.Cookie `json:"-"`
}

type Login struct {
	context.Context
}
