package session

import (
	"encoding/json"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	sessionList := SessionList{
		Context: r.Context(),
	}

	userSessions, err := sessionList.Execute()
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
