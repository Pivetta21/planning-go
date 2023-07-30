package auth

import (
	"context"

	"github.com/Pivetta21/planning-go/internal/repository"
)

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterOutput struct {
	Message string `json:"message"`
}

type Register struct {
	Context        context.Context
	UserRepository repository.IUserRepository
}

func NewRegister(ctx context.Context, userRepository *repository.UserRepository) *Register {
	return &Register{
		Context:        ctx,
		UserRepository: userRepository,
	}
}
