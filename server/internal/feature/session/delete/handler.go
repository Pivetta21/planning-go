package session

import (
	"net/http"

	"github.com/Pivetta21/planning-go/internal/util"
)

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	identifier, err := util.ExtractPathParameter(r, "identifier")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sessionDelete := SessionDelete{
		r.Context(),
	}

	if err := sessionDelete.Execute(identifier); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
}
