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

	feat := Delete{
		r.Context(),
	}

	if err := feat.Execute(identifier); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
}
