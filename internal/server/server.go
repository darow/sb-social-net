package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"sb_social_network/internal/store"
)

type server struct {
	router *chi.Mux
	store  store.Store
}

func NewServer(store store.Store) *server {
	s := &server{
		router: chi.NewRouter(),
		store:  store,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.MethodFunc("POST", "/create", s.handleCreate())
	s.router.MethodFunc("POST", "/make_friends", s.handleMakeFriends())
	s.router.MethodFunc("DELETE", "/user", s.handleDelete())
}
