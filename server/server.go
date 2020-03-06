package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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
	// Players are human unless p{n}ai=true in query
	rMatch.HandleFunc("/new", NewMatchHandler)
	rMatch.Path("/new").Queries("p1ai", "{p1ai:true|false}", "p2ai", "{p2ai:true|false}").HandlerFunc(NewMatchHandler)

	// Subrouter for /match/{id}
	rMatchId := rMatch.PathPrefix("/{id:[0-9]+}").Subrouter()
	rMatchId.HandleFunc("", GetMatchHandler)
	rMatchId.HandleFunc("/move", PostMoveHandler).Methods("POST")

	err := http.ListenAndServe(":4242", r)
	if err != nil {
		log.Fatal("Unable to run server")
	}
}
