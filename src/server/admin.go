package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"unknspec/src/models"
)

type AdminPage struct {
	layoutPath string
}

func NewAdminPage(layoutPath string) *AdminPage {
	return &AdminPage{layoutPath: layoutPath}
}

func (p *AdminPage) RenderPage() (*template.Template, error) {
	layout := `{{template "header"}}{{template "sidenav"}}{{template "dashboard"}}{{template "footer"}}`
	tmpl, err := template.New("layour").Parse(layout)
	if err != nil {
		return nil, err
	}
	return tmpl.ParseFiles(p.layoutPath)
}

func (p *AdminPage) Render(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(p.layoutPath)
	if err != nil {
		return nil, err
	}
	tmpl = tmpl.Lookup(name)
	if tmpl == nil {
		return nil, fmt.Errorf("no such template %s", name)
	}
	return tmpl, nil
}

func (s *Server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := s.adminPage.RenderPage()
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.Execute(w, nil)
}

func (s *Server) adminDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := s.adminPage.Render("dashboard")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.Execute(w, nil)
}

func (s *Server) adminDashboardTime(w http.ResponseWriter, r *http.Request) {
	tmpl, err := s.adminPage.Render("time")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.Execute(w, ResponseMap{"time": time.Now().Format(time.RFC822)})
}

func (s *Server) adminDashboardTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		data := r.PostFormValue("task")
		task := models.AdminTask{
			Task: data,
		}
		if err := s.db.CreateAdminTask(r.Context(), &task); err != nil {
			log.Fatalln(err)
		}
	case "DELETE":
		id := r.URL.Query()["id"][0]
		if err := s.db.DeleteAdminTask(r.Context(), id); err != nil {
			log.Fatalln(err)
		}
	}
	tmpl, err := s.adminPage.Render("tasks")
	if err != nil {
		log.Fatalln(err)
	}
	tasks, err := s.db.GetAdminTasks(r.Context())
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.Execute(w, ResponseMap{"tasks": tasks})
}

func (s *Server) adminArticles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := s.adminPage.Render("articles")
		if err != nil {
			log.Fatalln(err)
		}
		articles, err := s.db.GetAllArticles(r.Context())
		if err != nil {
			log.Fatalln(err)
		}
		tmpl.Execute(w, ResponseMap{"articles": articles})
	case "POST":
		tmpl, err := s.adminPage.Render("articlesTable")
		if err != nil {
			log.Fatalln(err)
		}
		var articles []*models.Article
		search := r.FormValue("search")
		if search == "" {
			articles, err = s.db.GetAllArticles(r.Context())
		} else {
			articles, err = s.db.GetArticlesByTitle(r.Context(), search)
		}
		if err != nil {
			log.Fatalln(err)
		}
		tmpl.Execute(w, ResponseMap{"articles": articles})
	}
}

func (s *Server) adminCreateArticle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := s.adminPage.Render("createArticle")
		if err != nil {
			log.Fatalln(err)
		}
		tmpl.Execute(w, nil)
	case "POST":
		var article models.Article
		article.Title = r.PostFormValue("title")
		article.Abstract = r.PostFormValue("abstract")
		article.Body = r.PostFormValue("body")
		article.Author = "SpectralJager"
		article.IsDraft = true
		if val := r.PostFormValue("isDraft"); val != "on" {
			article.IsDraft = false
		}
		err := s.db.CreateArticle(r.Context(), &article)
		if err != nil {
			log.Fatalln(err)
		}
		r.Method = "GET"
		s.adminArticles(w, r)
	}
}

func (s *Server) adminEditArticle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"][0]
	switch r.Method {
	case "GET":
		tmpl, err := s.adminPage.Render("editArticle")
		if err != nil {
			log.Fatalln(err)
		}
		article, err := s.db.GetArticle(r.Context(), id)
		if err != nil {
			log.Fatalln(err)
		}
		tmpl.Execute(w, ResponseMap{"article": article})
	case "POST":
		var article models.Article
		article.Title = r.PostFormValue("title")
		article.Abstract = r.PostFormValue("abstract")
		article.Body = r.PostFormValue("body")
		article.Author = "SpectralJager"
		article.IsDraft = true
		if val := r.PostFormValue("isDraft"); val != "on" {
			article.IsDraft = false
		}
		_, err := s.db.UpdateArticle(r.Context(), id, &article)
		if err != nil {
			log.Fatalln(err)
		}
		r.Method = "GET"
		s.adminArticles(w, r)
	case "DELETE":
		err := s.db.DeleteArticle(r.Context(), id)
		if err != nil {
			log.Fatalln(err)
		}
		r.Method = "GET"
		s.adminArticles(w, r)
	}
}
