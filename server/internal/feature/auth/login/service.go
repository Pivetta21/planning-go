package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/data/entity"
	"github.com/Pivetta21/planning-go/internal/data/enum"
	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (f *Login) Execute(in *LoginInput) (*LoginOutput, error) {
	obfuscatedErr := errors.New("please check your credentials")

	user, err := f.fetchUser(in.Username, in.Password)
	if err != nil {
		return nil, obfuscatedErr
	}

	userSession, err := f.persistUserSession(user.Id, in.Origin)
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

func (f *Login) fetchUser(username, password string) (*entity.User, error) {
	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	row := db.Ctx.Conn.QueryRowContext(
		queryCtx,
		`
		SELECT id, username, password, created_at
		FROM public.users
		WHERE username = $1
		`,
		strings.ToLower(username),
	)

	var user entity.User
	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}

	passwordMatches := f.checkPassword(password, user.Password)
	if !passwordMatches {
		return nil, errors.New("please check your credentials")
	}

	return &user, nil
}

func (f *Login) checkPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (f *Login) persistUserSession(userId int64, origin enum.SessionOrigin) (*entity.UserSession, error) {
	userSession, err := entity.NewUserSession(0, userId, core.CookieDurationAuthSession, uuid.New(), origin)
	if err != nil {
		return nil, err
	}

	queryCtx, cancel := context.WithTimeout(f.Context, db.Ctx.DefaultTimeout)
	defer cancel()

	row := db.Ctx.Conn.QueryRowContext(
		queryCtx,
		`
		INSERT INTO public.user_sessions(user_id, identifier, opaque_token, origin, expires_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
		`,
		userSession.UserId, userSession.Identifier, userSession.OpaqueToken, userSession.Origin, userSession.ExpiresAt,
	)

	if err := row.Scan(&userSession.Id); err != nil {
		return nil, err
	}

	return userSession, nil
}
