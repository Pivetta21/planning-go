package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/Pivetta21/planning-go/internal/data/enum"
	"github.com/Pivetta21/planning-go/internal/util"
	"github.com/google/uuid"
)

type UserSession struct {
	Id          int64
	UserId      int64
	Identifier  string
	ExpiresAt   time.Time
	OpaqueToken uuid.UUID
	Origin      enum.SessionOrigin
	CreatedAt   time.Time
}

func NewUserSession(
	id int64,
	userId int64,
	expiration time.Duration,
	opaqueToken uuid.UUID,
	origin enum.SessionOrigin,
) (*UserSession, error) {
	userSession := &UserSession{
		Id:          id,
		UserId:      userId,
		Identifier:  util.GenerateIdentifier(),
		ExpiresAt:   time.Now().UTC().Add(expiration),
		OpaqueToken: opaqueToken,
		Origin:      origin,
	}

	if err := userSession.Validate(); err != nil {
		return nil, err
	}

	return userSession, nil
}

func (e *UserSession) Validate() error {
	errs := make([]string, 0, 6)

	if e.UserId == 0 {
		errs = append(errs, "user id is required")
	}

	if e.Identifier == "" {
		errs = append(errs, "identifier is required")
	}

	if e.ExpiresAt.IsZero() {
		errs = append(errs, "expires at is required")
	}

	if e.OpaqueToken == uuid.Nil {
		errs = append(errs, "opaque token is required")
	}

	if e.Origin == 0 {
		errs = append(errs, "origin is required")
	}

	if ok := e.Origin.IsDefined(); !ok {
		errs = append(errs, "invalid value for origin")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}
