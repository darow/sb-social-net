package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sb_social_network/internal/model"
	"strconv"
)

func (s *server) Create() http.HandlerFunc {
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

func (s *server) MakeFriends() http.HandlerFunc {
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

		s.respond(w, r, http.StatusOK, map[string]string{"msg": fmt.Sprintf("%s добавил в друзья пользователя %s", u1.Name, u2.Name)})
	}
}

func (s *server) GetFriends() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ctxUserKey{}).(*model.User)

		if !ok {
			s.respondError(w, r, http.StatusInternalServerError, ErrCtxDoesNotExist)
		}

		size := len(user.Friends) * 4
		for _, v := range user.Friends {
			size += len(strconv.Itoa(v.ID)) + len(v.Name) + len(strconv.Itoa(v.Age))
		}

		buf := bytes.NewBuffer(make([]byte, 0, size))
		buf.WriteRune('[')
		for i, v := range user.Friends {
			if i != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("%s %s %s", strconv.Itoa(v.ID), v.Name, strconv.Itoa(v.Age)))
		}
		buf.WriteRune(']')

		s.respond(w, r, http.StatusOK, map[string]string{"friends_list": buf.String()})
	}
}

func (s *server) Delete() http.HandlerFunc {
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
