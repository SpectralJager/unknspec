package server

import (
	"log"
	"net/http"
	"unknspec/src/database"

	"github.com/gorilla/mux"
)

type ReponseMap map[string]any

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

/*
/ |>
	articles |>
		search
		update
	article
/admin |>
	articles |>
		search
	article |>
		new |>
			save
		edit |>
			save
			delete
	result
*/

func (s *Server) Run() {
	router := mux.NewRouter()

	log.Println("initialize server routes...")

	temp := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/public/").Handler(temp)

	router.Use(s.logger)

	// admin routes
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/", s.handleAdminDashboard).Methods("GET")
	adminRouter.HandleFunc("/dashboard/tasks", s.handleAdminDashboardTasks).Methods("DELETE", "POST")
	adminRouter.HandleFunc("/dashboard/cpu", s.handleAdminDashboardCpu).Methods("GET")

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
