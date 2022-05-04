package teststore

import (
	"errors"
	"sb_social_network/internal/model"
)

var (
	idCounter int
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(user model.User) int {
	idCounter++
	user.ID = idCounter
	r.users[idCounter] = &user
	return idCounter
}

func (r *UserRepository) MakeFriends(user1, user2 model.User) {
	if !r.contains(r.users[user1.ID].Friends, &user2) {
		r.users[user1.ID].Friends = append(r.users[user1.ID].Friends, &user2)
	}
}

func (r *UserRepository) contains(m []*model.User, e *model.User) bool {
	for _, a := range m {
		if a == e {
			return true
		}
	}

	return false
}

func (r *UserRepository) FindByID(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("error! FindByID user not found")
	}

	return u, nil
}
