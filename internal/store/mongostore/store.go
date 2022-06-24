package mongostore

import (
	"sb_social_network/internal/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db             *mongo.Database
	userRepository store.UserRepository
}

func New(db *mongo.Database) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	c := s.db.Collection("users")
	s.userRepository = &UserRepository{
		store: s,
		users: c,
	}

	return s.userRepository
}
