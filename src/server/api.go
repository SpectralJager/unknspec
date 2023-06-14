package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) apiEditArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	article, err := s.db.GetArticle(ctx, id)
	if err != nil {
		log.Println(err)
		return
	}
	article.Title = r.PostFormValue("title")
	article.Abstract = r.PostFormValue("abstract")
	if is_draft := r.PostFormValue("is_draft"); is_draft != "" {
		if is_draft == "on" {
			article.IsDraft = true
		}
	} else {
		article.IsDraft = false
	}
	s.db.UpdateArticle(ctx, id, article)
}
