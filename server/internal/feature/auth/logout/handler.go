package auth

import (
	"net/http"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	logout := Logout{
		Context: r.Context(),
	}

	out, err := logout.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	http.SetCookie(w, out.Cookie)

	w.WriteHeader(http.StatusOK)
}
