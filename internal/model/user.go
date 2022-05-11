package model

type User struct {
	ID      int
	Name    string
	Age     int
	Friends []*User
}

func (u *User) RemoveFromFriends(user *User) {
	for i, friend := range u.Friends {
		if friend == user {
			u.Friends[i], u.Friends[len(u.Friends)-1] = u.Friends[len(u.Friends)-1], u.Friends[i]
			u.Friends = u.Friends[:len(u.Friends)-1]

			break
		}
	}
}
