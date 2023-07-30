package auth

import (
	"net/http"

	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/Pivetta21/planning-go/internal/repository"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	userSessionRepository := repository.NewUserSessionRepository(db.Ctx)
	logout := NewLogout(r.Context(), userSessionRepository)

	out, err := logout.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	http.SetCookie(w, out.Cookie)

	w.WriteHeader(http.StatusOK)
}
