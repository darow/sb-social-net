package store

import "sb_social_network/internal/models"

type Store interface {
	User() UserRepository
}

type UserRepository interface {
	Create(models.User) string
	MakeFriends(u1, u2 models.User)
	Delete(*models.User)
	FindByID(string) (*models.User, error)
	SetAge(*models.User, int)
}
