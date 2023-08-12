package update

import (
	"encoding/json"
	"net/http"
)

func HandleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	var in Input
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "please verify the provided payload", http.StatusBadRequest)
		return
	}

	feat := Update{
		r.Context(),
	}

	if err := feat.Execute(in); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
