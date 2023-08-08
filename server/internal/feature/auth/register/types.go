package auth

import (
	"context"
)

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterOutput struct {
	Message string `json:"message"`
}

type Register struct {
	context.Context
}
