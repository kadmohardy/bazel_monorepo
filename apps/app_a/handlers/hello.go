package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	strUtil "github.com/kadmohardy/bazel_monorepo/lib/strings"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strUtil.ToUppercase(vars["name"])
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message on app a": fmt.Sprintf("Hello %s", name),
	})
}
