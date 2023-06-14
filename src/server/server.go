package server

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"
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

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/article/{id:[a-zA-Z0-9]+}/edit", s.apiEditArticle)

	admin := router.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/articles", s.adminArticles)
	admin.HandleFunc("/article/{id:[a-zA-Z0-9]+}", s.adminEditArticle)

	router.HandleFunc("/articles", s.renderArticles)
	router.HandleFunc("/article/{id:[a-zA-Z0-9]+}", s.renderArticle).Methods("GET")

	log.Printf("serve on %s ...\n", s.addr)
	http.ListenAndServe(s.addr, router)
}

func (s *Server) adminArticles(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	articles, err := s.db.GetAllArticles(ctx)
	if err != nil {
		log.Println(err)
	}
	tmpl, err := template.ParseFiles("templates/admin/admin.html")
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, ReponseMap{"articles": articles})
}

func (s *Server) adminEditArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Println("cant extract 'id'")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	article, err := s.db.GetArticle(ctx, id)
	if err != nil {
		log.Println(err)
	}
	tmpl, err := template.ParseFiles("templates/admin/article.html")
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, ReponseMap{"article": article})
}
