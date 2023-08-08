package auth

import (
	"context"
	"net/http"
)

type RefreshOutput struct {
	Cookie *http.Cookie `json:"-"`
}

type Refresh struct {
	context.Context
}
