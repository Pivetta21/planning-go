package auth

import (
	"encoding/json"
	"net/http"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var in RegisterInput
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "please verify the provided payload", http.StatusBadRequest)
		return
	}

	register := Register{
		Context: r.Context(),
	}

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
