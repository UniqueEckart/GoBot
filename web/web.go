package web

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func bothealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bot is healthy!")
}

func Init() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/healthcheck"), bothealth)

	// Only for Debug
	fmt.Printf("[LOG] HTTP Multiplexer listening\n")
	go http.ListenAndServe("localhost:8080", mux)
}
