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
		s.checkErr(w, r, http.StatusBadRequest, err)

		user, err := s.store.User().FindByID(userID)
		s.checkErr(w, r, http.StatusBadRequest, err)

		ctx := context.WithValue(r.Context(), ctxUserKey{}, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
