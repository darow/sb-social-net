package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (s *server) userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIDParam := chi.URLParam(r, "userID")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
			return
		}

		user, err := s.store.User().FindByID(userID)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), ctxUserKey{}, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
