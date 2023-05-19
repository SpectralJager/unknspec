package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", basicHandle(indexHandler))

	http.ListenAndServe(":8080", router)
}

type handleFunc func(http.ResponseWriter, *http.Request) error

func basicHandle(fn handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return err
	}
	return tmpl.Execute(w, nil)
}
