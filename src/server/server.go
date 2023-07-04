package server

import (
	"html/template"
	"log"
	"net/http"
	"unknspec/src/database"

	"github.com/gorilla/mux"
)

type ResponseMap map[string]any

type Page interface {
	RenderPage() (*template.Template, error)
	Render(name string) (*template.Template, error)
}

type Server struct {
	addr      string
	db        *database.MongoStorage
	adminPage Page
	mainPage  Page
}

func NewServer(addr string, db *database.MongoStorage) *Server {
	return &Server{
		addr:      addr,
		db:        db,
		adminPage: NewAdminPage("templates/admin.html"),
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	log.Println("initialize server routes...")

	temp := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/public/").Handler(temp)

	router.Use(s.logger)

	// admin routes
	admin := router.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/", s.handleAdmin).Methods("GET")
	admin.HandleFunc("/dashboard", s.adminDashboard).Methods("GET")
	admin.HandleFunc("/dashboard/time", s.adminDashboardTime).Methods("GET")
	admin.HandleFunc("/dashboard/tasks", s.adminDashboardTasks).Methods("GET", "POST", "DELETE")
	admin.HandleFunc("/articles", s.adminArticles).Methods("GET", "POST")
	admin.HandleFunc("/article/new", s.adminCreateArticle).Methods("GET", "POST")
	admin.HandleFunc("/article/edit", s.adminEditArticle).Methods("GET", "POST", "DELETE")
	// public routes

	log.Printf("serve on %s ...\n", s.addr)
	http.ListenAndServe(s.addr, router)
}

func (s *Server) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s by %s\n", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
