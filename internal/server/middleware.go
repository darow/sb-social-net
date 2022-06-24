package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *server) userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")

		user, err := s.store.User().FindByID(userID)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), ctxUserKey{}, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
