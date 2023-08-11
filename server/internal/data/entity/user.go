package entity

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	Id           int64
	Username     string
	Password     string
	CreatedAt    time.Time
	SessionLimit int
}

func NewUser(id int64, username string, password string) (*User, error) {
	user := &User{
		Id:       id,
		Username: strings.ToLower(username),
		Password: password,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (e *User) Validate() error {
	errs := make([]string, 0, 2)

	if e.Username == "" {
		errs = append(errs, "username is required")
	}

	if e.Password == "" {
		errs = append(errs, "opaque token is required")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}
