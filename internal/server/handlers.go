package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sb_social_network/internal/model"
	"strconv"
)

func (s *server) handleCreate() http.HandlerFunc {
	type request struct {
		Name    string `json:"name"`
		Age     string `json:"age"`
		Friends []int  `json:"friends"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
		}

		age, err := strconv.Atoi(req.Age)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
		}

		users := make([]*model.User, 0)
		for _, v := range req.Friends {
			u, err := s.store.User().FindByID(v)
			if err != nil {
				s.respondError(w, r, http.StatusBadRequest, err)
			}
			users = append(users, u)
		}

		u := model.User{
			Name:    req.Name,
			Age:     age,
			Friends: users,
		}
		id := s.store.User().Create(u)
		s.respond(w, r, http.StatusCreated, map[string]int{"id": id})
	}
}

func (s *server) handleMakeFriends() http.HandlerFunc {
	type request struct {
		SourceID int `json:"source_id"`
		TargetID int `json:"target_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
		}

		u1, err := s.store.User().FindByID(req.SourceID)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
		}
		u2, err := s.store.User().FindByID(req.TargetID)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
		}

		s.store.User().MakeFriends(*u1, *u2)

		s.respond(w, r, http.StatusOK, map[string]string{"msg": fmt.Sprintf("%s и %s теперь друзья", u1.Name, u2.Name)})
	}
}

func (s *server) handleDelete() http.HandlerFunc {
	type request struct {
		TargetID int `json:"target_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
		}

		u1, err := s.store.User().FindByID(req.TargetID)
		if err != nil {
			s.respondError(w, r, http.StatusBadRequest, err)
		}

		s.store.User().Delete(u1)

		s.respond(w, r, http.StatusOK, map[string]string{"msg": fmt.Sprintf("%s удален", u1.Name)})
	}
}
