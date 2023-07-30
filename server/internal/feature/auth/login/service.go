package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/data/entity"
	"github.com/Pivetta21/planning-go/internal/data/enum"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *Login) Execute(in *LoginInput) (*LoginOutput, error) {
	obfuscatedErr := errors.New("please check your credentials")

	user, err := u.fetchUser(in.Username, in.Password)
	if err != nil {
		return nil, obfuscatedErr
	}

	userSession, err := u.persistUserSession(user.Id, in.Origin)
	if err != nil {
		return nil, obfuscatedErr
	}

	cookie := &http.Cookie{
		Name:     core.CookieNameAuthSession.String(),
		Value:    userSession.OpaqueToken.String(),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   int(core.CookieDurationAuthSession.Seconds()),
	}

	out := &LoginOutput{
		Message: fmt.Sprintf("Welcome back %#q!", user.Username),
		Cookie:  cookie,
	}

	return out, nil
}

func (u *Login) fetchUser(username, password string) (*entity.User, error) {
	user, err := u.UserRepository.GetByUsername(u.Context, username)
	if err != nil {
		return nil, err
	}

	passwordMatches := u.checkPassword(password, user.Password)
	if !passwordMatches {
		return nil, err
	}

	return user, nil
}

func (u *Login) checkPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (u *Login) persistUserSession(userId int64, origin enum.SessionOrigin) (*entity.UserSession, error) {
	userSession, err := entity.NewUserSession(0, userId, core.CookieDurationAuthSession, uuid.New(), origin)
	if err != nil {
		return nil, err
	}

	id, err := u.UserSessionRepository.Save(u.Context, userSession)
	if err != nil {
		return nil, err
	}

	userSession.Id = id

	return userSession, nil
}
