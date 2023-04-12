package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	s := CreateNewServer()
	s.MountHandlers()
	err := http.ListenAndServe(":3333", s.Router)
	if err != nil {
		return
	}
}

type Server struct {
	Router *chi.Mux
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {
	s.Router.Use(middleware.Logger)

	s.Router.Get("/", HelloWorld)
}
