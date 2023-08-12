package auth

import (
	"net/http"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	feat := Logout{
		Context: r.Context(),
	}

	out, err := feat.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	http.SetCookie(w, out.Cookie)

	w.WriteHeader(http.StatusOK)
}
