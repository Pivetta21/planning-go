package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/Pivetta21/planning-go/internal/repository"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var in RegisterInput
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "please verify the provided payload", http.StatusBadRequest)
		return
	}

	userRepository := repository.NewUserRepository(db.Ctx)
	register := NewRegister(r.Context(), userRepository)

	out, err := register.Execute(&in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if json.NewEncoder(w).Encode(out) != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}
