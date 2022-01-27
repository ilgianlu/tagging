package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to develop - testing"))
	})
	r.Get("/staging", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to staging"))
	})
	r.Get("/production", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to production"))
	})
	log.Println("listening", os.Getenv("LISTEN_INTERFACE"))
	http.ListenAndServe(os.Getenv("LISTEN_INTERFACE"), r)
}
