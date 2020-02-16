package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	log.Println("Starting Gomoku in server mode")

	// Basic router
	r := mux.NewRouter()

	r.Use(corsMiddleware)
	r.HandleFunc("/", HomeHandler).Methods("GET", "POST")
	r.Use(mux.CORSMethodMiddleware(r))
	http.Handle("/", r)

	// Subrouter for /match
	rMatch := r.PathPrefix("/match").Subrouter()
	rMatch.HandleFunc("/new", NewMatchHandler)

	// Subrouter for /match/{id}
	rMatchId := rMatch.PathPrefix("/{id:[0-9]+}").Subrouter()
	rMatchId.HandleFunc("", GetMatchHandler)
	rMatchId.HandleFunc("/move", PostMoveHandler).Methods("POST")

	err := http.ListenAndServe(":4242", r)
	if err != nil {
		log.Fatal("Unable to run server")
	}
}
