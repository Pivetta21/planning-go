package auth

import (
	"context"
)

type Input struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Output struct {
	Message string `json:"message"`
}

type Register struct {
	context.Context
}
