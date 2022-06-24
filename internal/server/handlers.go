package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"sb_social_network/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *server) Create() http.HandlerFunc {
	type request struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Friends []string `json:"friends"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
			return
		}

		friends := make([]primitive.ObjectID, 0)
		for _, v := range req.Friends {
			u, err := s.store.User().FindByID(v)
			if err != nil {
				s.respondError(w, r, http.StatusBadRequest, err)
				return
			}
			friends = append(friends, u.ID)
		}

		u := models.User{
			Name:    req.Name,
			Age:     req.Age,
			Friends: friends,
		}
		id := s.store.User().Create(u)
		s.respond(w, r, http.StatusCreated, map[string]any{"id": id})
	}
}

func (s *server) MakeFriends() http.HandlerFunc {
	type request struct {
		SourceID string `json:"source_id"`
		TargetID string `json:"target_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
			return
		}

		u1, err := s.store.User().FindByID(req.SourceID)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
			return
		}
		u2, err := s.store.User().FindByID(req.TargetID)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
			return
		}

		s.store.User().MakeFriends(*u1, *u2)

		s.respond(w, r, http.StatusOK, map[string]string{"msg": fmt.Sprintf("%s и %s теперь друзья", u1.Name, u2.Name)})
	}
}

func (s *server) Delete() http.HandlerFunc {
	type request struct {
		TargetID string `json:"target_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
			return
		}

		u1, err := s.store.User().FindByID(req.TargetID)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
			return
		}

		s.store.User().Delete(u1)

		s.respond(w, r, http.StatusOK, map[string]string{"msg": fmt.Sprintf("%s удален", u1.Name)})
	}
}

func (s *server) GetFriends() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ctxUserKey{}).(*models.User)

		if !ok {
			s.respondError(w, r, http.StatusInternalServerError, ErrCtxDoesNotExist)
			return
		}

		size := len(user.Friends) * 4
		for _, v := range user.Friends {
			size += len(fmt.Sprintf("%s", v))
		}

		buf := bytes.NewBuffer(make([]byte, 0, size))
		buf.WriteRune('[')
		for i, v := range user.Friends {
			if i != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("%s", v))
		}
		buf.WriteRune(']')

		s.respond(w, r, http.StatusOK, map[string]string{"friends_list": buf.String()})
	}
}

func (s *server) SetAge() http.HandlerFunc {
	type request struct {
		NewAge int `json:"new age"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		user, ok := r.Context().Value(ctxUserKey{}).(*models.User)

		if !ok {
			s.respondError(w, r, http.StatusInternalServerError, ErrCtxDoesNotExist)
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
		}

		previousAge := user.Age
		s.store.User().SetAge(user, req.NewAge)

		s.respond(w, r, http.StatusOK, map[string]string{"msg": fmt.Sprintf("возраст %s изменен с %d на %d", user.Name, previousAge, req.NewAge)})
	}
}
