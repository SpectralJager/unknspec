package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
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
	tmpl.Execute(w, ReponseMap{"time": time.Now().Format(time.RFC822)})
}

func (s *Server) adminArticles(w http.ResponseWriter, r *http.Request) {
	tmpl, err := s.adminPage.Render("articles")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.Execute(w, nil)
}
