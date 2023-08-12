package update

import (
	"net/http"
)

func HandleDeleteProfile(w http.ResponseWriter, r *http.Request) {
	feat := Delete{
		r.Context(),
	}

	if err := feat.Execute(); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
