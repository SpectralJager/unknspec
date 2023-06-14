package server

import (
	"log"
	"net/http"
	"unknspec/src/database"

	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	db   *database.MongoStorage
}

func NewServer(addr string, db *database.MongoStorage) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	log.Println("initialize server routes...")

	temp := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/public/").Handler(temp)

	// api := router.PathPrefix("/api").Subrouter()

	router.HandleFunc("/articles", s.renderArticles)
	router.HandleFunc("/article/{id:[a-zA-Z0-9]+}", s.renderArticle).Methods("GET")

	log.Printf("serve on %s ...\n", s.addr)
	http.ListenAndServe(s.addr, router)
}
