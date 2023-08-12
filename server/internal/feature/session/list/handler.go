package session

import (
	"encoding/json"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	feat := List{
		Context: r.Context(),
	}

	out, err := feat.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if json.NewEncoder(w).Encode(out) != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}
