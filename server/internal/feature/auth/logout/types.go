package auth

import (
	"context"
	"net/http"
)

type LogoutOutput struct {
	Cookie *http.Cookie `json:"-"`
}

type Logout struct {
	context.Context
}
