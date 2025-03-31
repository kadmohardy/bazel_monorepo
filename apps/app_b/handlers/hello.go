package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message on app b": fmt.Sprintf("Hello %s", vars["name"]),
	})
}
