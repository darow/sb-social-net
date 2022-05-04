package store

import "sb_social_network/internal/model"

type Store interface {
	User() UserRepository
}

type UserRepository interface {
	Create(model.User) int
	MakeFriends(u1, u2 model.User)
	FindByID(int) (*model.User, error)
}
