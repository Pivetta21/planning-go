package session

import (
	"encoding/json"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	userSessions, err := ExecuteList(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if json.NewEncoder(w).Encode(userSessions) != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
}
