package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
	"unknspec/src/models"

	"github.com/gorilla/mux"
)

func (s *Server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/admin.html")
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) recieveError(w http.ResponseWriter, e error) {
	tmpl, err := template.ParseFiles("templates/admin/error.html")
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl.ExecuteTemplate(w, "error", ReponseMap{"msg": e.Error()}); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) recieveArticles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		articles, err := s.db.GetAllArticles(ctx)
		if err != nil {
			s.recieveError(w, err)
			return
		}
		tmpl, err := template.ParseFiles("templates/admin/articles.html")
		if err != nil {
			s.recieveError(w, err)
			return
		}
		if err := tmpl.ExecuteTemplate(w, "articles", ReponseMap{"articles": articles}); err != nil {
			s.recieveError(w, err)
			return
		}
	case "POST":
		r.ParseForm()
		search := r.Form.Get("search")
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		articles, err := s.db.GetArticlesByTitle(ctx, search)
		if err != nil {
			s.recieveError(w, err)
			return
		}
		tmpl, err := template.ParseFiles("templates/admin/articles.html")
		if err != nil {
			s.recieveError(w, err)
			return
		}
		if err := tmpl.ExecuteTemplate(w, "articlesList", ReponseMap{"articles": articles}); err != nil {
			s.recieveError(w, err)
			return
		}

	default:
		s.recieveError(w, fmt.Errorf("unsupported method %s for URL:%s", r.Method, r.RequestURI))
	}
}

func (s *Server) recieveArticle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/admin/article.html")
		if err != nil {
			s.recieveError(w, err)
			return
		}
		if err := tmpl.ExecuteTemplate(w, "newArticle", nil); err != nil {
			s.recieveError(w, err)
			return
		}
	case "POST":
		article := models.Article{}
		r.ParseForm()
		article.Title = r.Form.Get("title")
		article.Abstract = r.Form.Get("abstract")
		article.Body = r.Form.Get("body")
		article.IsDraft = r.Form.Get("isDraft") == "on"
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		err := s.db.CreateArticle(ctx, &article)
		if err != nil {
			s.recieveError(w, err)
			return
		}
		r.Method = "GET"
		s.recieveArticles(w, r)
	default:
		s.recieveError(w, fmt.Errorf("unsupported method %s for URL:%s", r.Method, r.RequestURI))
	}
}

func (s *Server) recieveArticleId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	switch r.Method {
	case "GET":
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		article, err := s.db.GetArticle(ctx, id)
		if err != nil {
			s.recieveError(w, err)
			return
		}
		tmpl, err := template.ParseFiles("templates/admin/article.html")
		if err != nil {
			s.recieveError(w, err)
			return
		}
		if err := tmpl.ExecuteTemplate(w, "editArticle", ReponseMap{"article": article}); err != nil {
			s.recieveError(w, err)
			return
		}
	case "PUT":
		article := models.Article{}
		r.ParseForm()
		article.Title = r.Form.Get("title")
		article.Abstract = r.Form.Get("abstract")
		article.Body = r.Form.Get("body")
		article.IsDraft = r.Form.Get("isDraft") == "on"
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		_, err := s.db.UpdateArticle(ctx, id, &article)
		if err != nil {
			s.recieveError(w, err)
			return
		}
		r.Method = "GET"
		s.recieveArticles(w, r)
	case "DELETE":
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		err := s.db.DeleteArticle(ctx, id)
		if err != nil {
			s.recieveError(w, err)
			return
		}
		r.Method = "GET"
		s.recieveArticles(w, r)
	default:
		s.recieveError(w, fmt.Errorf("unsupported method %s for URL:%s", r.Method, r.RequestURI))
	}
}
