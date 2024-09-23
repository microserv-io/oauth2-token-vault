package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	r := chi.NewRouter()

	r.Get("/oauth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oauth"))
	})

	r.Get("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oauth token"))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
