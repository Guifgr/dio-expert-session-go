package actuator

import (
	"encoding/json"
	"net/http"
)

type HeathBody struct {
	Status string `json:"status"`
}

//Heath documentation
func Heath(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transaction = HeathBody{Status: "alive"}
	_ = json.NewEncoder(w).Encode(transaction)
}
