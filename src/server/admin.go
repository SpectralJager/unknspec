package server

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"text/template"
	"time"
)

const (
	adminRootPath      = "templates/admin/"
	adminBaseTemplate  = adminRootPath + "base.html"
	adminScreensPath   = adminRootPath + "screens/"
	adminFragmentsPath = adminRootPath + "fragments/"
)

var (
	taskDb = []string{
		"Some permanent task",
	}
)

func (s *Server) handleAdminDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(adminScreensPath+"dashboard.html", adminBaseTemplate, adminFragmentsPath+"sidenav.html")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.Execute(w, ReponseMap{
		"tasks": taskDb,
		"time":  getTime(),
	})
}

func (s *Server) handleAdminDashboardTasks(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(adminScreensPath + "dashboard.html")
	if err != nil {
		log.Fatalln(err)
	}
	switch r.Method {
	case "POST":
		r.ParseForm()
		task := r.FormValue("todo")
		taskDb = append(taskDb, task)
	case "DELETE":
		params, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Fatalln(err)
		}
		id, err := strconv.ParseInt(params.Get("id"), 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		taskDb = append(taskDb[:id], taskDb[id+1:]...)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tmpl.ExecuteTemplate(w, "tasks", ReponseMap{
		"tasks": taskDb,
	})
}

func (s *Server) handleAdminDashboardTime(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(adminScreensPath + "dashboard.html")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.ExecuteTemplate(w, "time", ReponseMap{
		"time": getTime(),
	})
}

func getTime() string {
	return time.Now().Format(time.RFC822)
}
