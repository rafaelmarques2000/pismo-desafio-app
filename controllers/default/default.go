package _default

import (
	"encoding/json"
	"net/http"
)

func DefaultIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string {
		"status":"OK",
	})
}
