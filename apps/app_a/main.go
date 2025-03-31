package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kadmohardy/bazel_monorepo/apps/app_a/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{name}", handlers.HandleHello).Methods("GET")
	print("Server started app_a...")

	http.ListenAndServe(":8080", r)
}
