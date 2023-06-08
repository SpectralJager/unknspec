package server

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
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

	temp := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/public/").Handler(temp)

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/articles", handleFuncWrapper(s.apiArticles)).Methods("GET")
	api.HandleFunc("/articles/{tag:[a-zA-Z0-9 ]+}", handleFuncWrapper(s.apiTag)).Methods("GET")
	api.HandleFunc("/article", handleFuncWrapper(s.apiArticle)).Methods("POST")
	api.HandleFunc("/article/{id:[a-zA-Z0-9]+}", handleFuncWrapper(s.apiArticle)).Methods("GET", "PUT", "DELETE")

	router.HandleFunc("/articles", s.renderArticles).Methods("GET")
	router.HandleFunc("/articles/{tag:[a-zA-Z0-9 ]+}", s.renderTag).Methods("GET")
	router.HandleFunc("/article/{id:[a-zA-Z0-9]+}", s.renderArticle).Methods("GET")

	log.Printf("serve on %s...\n", s.addr)
	http.ListenAndServe(s.addr, router)
}

func (s *Server) apiArticles(w http.ResponseWriter, r *http.Request) error {
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

func (s *Server) apiArticle(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	switch r.Method {
	case "GET":
		if id, ok := vars["id"]; ok {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			article, err := s.db.GetArticle(ctx, id)
			if err != nil {
				return err
			}
			return WriteJson(w, http.StatusOK, article)
		}
		return fmt.Errorf("expected id, got %v", vars)
	case "PUT":
		if id, ok := vars["id"]; ok {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			article := &models.Article{}
			json.NewDecoder(r.Body).Decode(article)
			article, err := s.db.UpdateArticle(ctx, id, article)
			if err != nil {
				return err
			}
			return WriteJson(w, http.StatusOK, article)
		}
		return fmt.Errorf("expected id, got %v", vars)
	case "DELETE":
		if id, ok := vars["id"]; ok {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err := s.db.DeleteArticle(ctx, id)
			if err != nil {
				return err
			}
			return WriteJson(w, http.StatusOK, ReponseMap{"result": "success", "id": id})
		}
		return fmt.Errorf("expected id, got %v", vars)

	case "POST":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		article := models.Article{}
		json.NewDecoder(r.Body).Decode(&article)
		return s.db.CreateArticle(ctx, &article)
	}
	return nil
}

func (s *Server) apiTag(w http.ResponseWriter, r *http.Request) error {
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

func (s *Server) renderArticles(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	articles, err := s.db.GetOnlyPublishedArticles(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	resp := ReponseMap{"articles": articles, "tags": map[string]uint8{}}
	for _, article := range articles {
		tags := article.Tags
		if tags != nil {
			for _, tag := range tags {
				resp["tags"].(map[string]uint8)[tag] = 0
			}
		}
	}
	tmpl, err := template.ParseFiles("templates/articles.html", "templates/base.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "articles", resp)
	if err != nil {
		log.Println(err)
		return
	}
}
func (s *Server) renderArticle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Printf("url %s expected id", r.URL.Path)
		return
	}
	articles, err := s.db.GetArticle(ctx, id)
	if err != nil {
		log.Println(err)
		return
	}
	resp := ReponseMap{"article": articles}
	tmpl, err := template.ParseFiles("templates/article.html", "templates/base.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "article", resp)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Server) renderTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag, ok := vars["tag"]
	if !ok {
		log.Printf("url %s expected tag", r.URL.Path)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	articles, err := s.db.GetOnlyPublishedArticlesWithTag(ctx, tag)
	if err != nil {
		log.Println(err)
		return
	}
	resp := ReponseMap{"articles": articles, "tag": tag}
	tmpl, err := template.ParseFiles("templates/tag.html", "templates/base.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "tag", resp)
	if err != nil {
		log.Println(err)
		return
	}
}
