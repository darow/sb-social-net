package teststore

import (
	"fmt"

	"sb_social_network/internal/store"
)

var (
	idCounter int
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(user model.User) string {
	idCounter++
	user.ID = fmt.Sprintf("%d", idCounter)
	r.users[user.ID] = &user

	return user.ID
}

func (r *UserRepository) MakeFriends(user1, user2 model.User) {
	if !r.contains(r.users[user1.ID].Friends, user2.ID) {
		r.users[user1.ID].Friends = append(r.users[user1.ID].Friends, user2.ID)
		r.users[user2.ID].Friends = append(r.users[user2.ID].Friends, user1.ID)
	}
}

// Delete Удаляем пользователя и ссылки на дружбу с ним у всех пользователей.
func (r *UserRepository) Delete(user *model.User) {
	for _, u := range r.users {
		u.RemoveFromFriends(user)
	}
	delete(r.users, user.ID)
}

func (r *UserRepository) SetAge(user *model.User, newAge int) error {
	u, ok := r.users[user.ID]
	if !ok {
		return store.ErrObjectNotFound
	}
	u.Age = newAge

	return nil
}

func (r *UserRepository) contains(m []string, id string) bool {
	for _, e := range m {
		if e == id {
			return true
		}
	}

	return false
}

func (r *UserRepository) FindByID(id string) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrObjectNotFound
	}

	return u, nil
}
