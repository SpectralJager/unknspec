package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"unknspec/src/database"
	"unknspec/src/models"

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
	router.HandleFunc("/posts", handleFuncWrapper(s.handleArticles))
	router.HandleFunc("/posts/{tag:[a-zA-Z0-9 ]+}", handleFuncWrapper(s.handleTag))
	router.HandleFunc("/post", handleFuncWrapper(s.handleArticle))

	log.Printf("serve on %s...\n", s.addr)
	http.ListenAndServe(s.addr, router)
}

func (s *Server) handleArticles(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("url: %s accepts only GET", r.URL.Path)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	articles, err := s.db.GetAllArticles(ctx)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, articles)
}

func (s *Server) handleArticle(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
	case "POST":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		article := models.Article{}
		json.NewDecoder(r.Body).Decode(&article)
		return s.db.CreateArticle(ctx, &article)
	}
	return nil
}

func (s *Server) handleTag(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("url: %s accepts only GET", r.URL.Path)
	}
	vars := mux.Vars(r)
	tag, ok := vars["tag"]
	if !ok {
		return fmt.Errorf("url %s expected tag", r.URL.Path)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	articles, err := s.db.GetArticlesWithTag(ctx, tag)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, articles)
}
