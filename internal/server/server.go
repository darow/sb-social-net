package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"sb_social_network/internal/store"
)

type ctxUserKey struct{}

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
	s.router.MethodFunc("POST", "/create", s.Create())
	s.router.MethodFunc("POST", "/make_friends", s.MakeFriends())
	s.router.MethodFunc("DELETE", "/user", s.Delete())
	s.router.Route("/friends/{userID}", func(r chi.Router) {
		r.Use(s.userCtx)
		r.Get("/", s.GetFriends())
	})
}
