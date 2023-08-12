package auth

import (
	"net/http"

	"github.com/Pivetta21/planning-go/internal/core"
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

	feat := Refresh{
		Context: r.Context(),
	}

	out, err := feat.Execute(opaqueToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, out.Cookie)

	w.WriteHeader(http.StatusOK)
}
