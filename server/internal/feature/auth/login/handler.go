package auth

import (
	"encoding/json"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var in LoginInput
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "please verify the provided payload", http.StatusBadRequest)
		return
	}

	login := Login{
		Context: r.Context(),
	}

	out, err := login.Execute(&in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	http.SetCookie(w, out.Cookie)

	w.Header().Add("Content-Type", "application/json")
	if json.NewEncoder(w).Encode(out) != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}
