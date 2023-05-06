package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type Server struct {
	ListenedAddr string
	Store        Storage
}

func NewServer(listenAddr string, storage Storage) *Server {
	return &Server{
		ListenedAddr: listenAddr,
		Store:        storage,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()

	// routes for api
	router.HandleFunc("/api/v1/posts", apiWrapper(s.apiPostsHandler))
	router.HandleFunc("/api/v1/posts/{id}", apiWrapper(s.apiPostHandler))

	log.Printf("start listening on %s ...", s.ListenedAddr)
	return http.ListenAndServe(s.ListenedAddr, router)
}

// api/v1 handlers

func (s *Server) apiPostsHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		post := new(Post)
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			return fmt.Errorf("can't decode request body: %s", err.Error())
		}
		post, err = s.Store.CreatePost(post)
		if err != nil {
			return err
		}
		writeJson(w, http.StatusOK, post)
		return nil
	} else if r.Method == "GET" {
		posts, err := s.Store.GetPosts()
		if err != nil {
			return err
		}
		writeJson(w, http.StatusOK, posts)
		return nil
	}
	return fmt.Errorf("unsupported method: %s", r.Method)
}

func (s *Server) apiPostHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Utils
type ResponseError struct {
	Error string `json:"error"`
}

func apiWrapper(fn apiFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJson(w, http.StatusBadRequest, ResponseError{err.Error()})
		}
	}
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
