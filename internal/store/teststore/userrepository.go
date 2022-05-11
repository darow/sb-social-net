package teststore

import (
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

// Delete Удаляем у всех пользователей из списка друзей. Можно переделать так, чтоб удалялось только у тех пользователей,
// котороые в списке друзей удаляемого.
func (r *UserRepository) Delete(user *model.User) {
	delete(r.users, user.ID)
	for _, u := range r.users {
		u.RemoveFromFriends(user)
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
		return nil, ErrObjectNotFound
	}

	return u, nil
}
