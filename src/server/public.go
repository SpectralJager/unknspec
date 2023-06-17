package server

import (
	"context"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
)

func (s *Server) handleArticles(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	articles, err := s.db.GetOnlyPublishedArticles(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	tmpl, err := template.ParseFiles("templates/articles.html", "templates/base.html")
	if err != nil {
		log.Fatalln(err)
	}
	if err := tmpl.Execute(w, ReponseMap{"articles": articles}); err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) handleArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	article, err := s.db.GetArticle(ctx, id)
	if err != nil {
		log.Fatalln(err)
	}
	tmpl, err := template.ParseFiles("templates/article.html", "templates/base.html")
	if err != nil {
		log.Fatalln(err)
	}
	body := string(markdown.ToHTML([]byte(article.Body), nil, nil))
	if err := tmpl.Execute(w, ReponseMap{"article": article, "body": body}); err != nil {
		log.Fatalln(err)
	}
}
