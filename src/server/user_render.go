package server

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) renderArticles(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	articles, err := s.db.GetOnlyPublishedArticles(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	resp := ReponseMap{"articles": articles}
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
