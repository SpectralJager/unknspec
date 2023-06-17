package server

import (
	"log"
	"net/http"
	"unknspec/src/database"

	"github.com/gorilla/mux"
)

type ReponseMap map[string]any

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

/*
/ |>
	articles |>
		search
		update
	article
/admin |>
	articles |>
		search
	article |>
		new |>
			save
		edit |>
			save
			delete
	result
*/

func (s *Server) Run() {
	router := mux.NewRouter()

	log.Println("initialize server routes...")

	temp := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/public/").Handler(temp)

	router.Use(s.logger)

	// admin routes
	admin := router.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/", s.handleAdmin)
	admin.HandleFunc("/articles", s.recieveArticles).Methods("GET", "POST")
	admin.HandleFunc("/article", s.recieveArticle).Methods("GET", "POST")
	admin.HandleFunc("/article/{id:[0-9a-zA-Z]+}", s.recieveArticleId).Methods("GET", "PUT", "DELETE")

	// public routes
	router.HandleFunc("/articles", s.handleArticles).Methods("GET", "POST")
	router.HandleFunc("/article/{id:[0-9a-zA-Z]+}", s.handleArticle).Methods("GET")

	log.Printf("serve on %s ...\n", s.addr)
	http.ListenAndServe(s.addr, router)
}

func (s *Server) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s by %s\n", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
