package profile

import (
	"encoding/json"
	"net/http"
)

func HandleFindProfile(w http.ResponseWriter, r *http.Request) {
	feat := Find{
		r.Context(),
	}

	user, err := feat.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}
