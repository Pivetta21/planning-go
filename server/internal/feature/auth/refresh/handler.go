package auth

import (
	"net/http"

	"github.com/Pivetta21/planning-go/internal/core"
	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/Pivetta21/planning-go/internal/repository"
	"github.com/google/uuid"
)

func HandleRefresh(w http.ResponseWriter, r *http.Request) {
	authSessionCookie, err := r.Cookie(core.CookieNameAuthSession.String())
	if err != nil || authSessionCookie.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	opaqueToken, err := uuid.Parse(authSessionCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userSessionRepository := repository.NewUserSessionRepository(db.Ctx)
	refresh := NewRefresh(r.Context(), userSessionRepository)

	out, err := refresh.Execute(opaqueToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, out.Cookie)

	w.WriteHeader(http.StatusOK)
}
