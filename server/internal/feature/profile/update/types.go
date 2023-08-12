package update

import (
	"context"
)

type Input struct {
	Username string `json:"username"`
}

type Update struct {
	context.Context
}
