package auth

import (
	"fmt"

	"github.com/Pivetta21/planning-go/internal/data/entity"
	"golang.org/x/crypto/bcrypt"
)

func (u *Register) Execute(in *RegisterInput) (*RegisterOutput, error) {
	hashedPassword, err := u.hashPassword(in.Password)
	if err != nil {
		return nil, err
	}

	user, err := entity.NewUser(0, in.Username, hashedPassword)
	if err != nil {
		return nil, err
	}

	_, err = u.UserRepository.Save(u.Context, user)
	if err != nil {
		return nil, err
	}

	out := &RegisterOutput{
		Message: fmt.Sprintf("User %#q created successfully", in.Username),
	}

	return out, err
}

func (u *Register) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}
