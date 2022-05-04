package teststore

import (
	"sb_social_network/internal/model"
	"sb_social_network/internal/store"
)

type Store struct {
	userRepository store.UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		users: make(map[int]*model.User, 0),
	}

	return s.userRepository
}
