package auth

import (
	"context"
	"net/http"
)

type Output struct {
	Cookie *http.Cookie `json:"-"`
}

type Refresh struct {
	context.Context
}
