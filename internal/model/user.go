package model

type User struct {
	ID      string `bson:"id,omitempty"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []string
}

func (u *User) RemoveFromFriends(user *User) {
	for i, friendID := range u.Friends {
		if friendID == user.ID {
			u.Friends[i], u.Friends[len(u.Friends)-1] = u.Friends[len(u.Friends)-1], u.Friends[i]
			u.Friends = u.Friends[:len(u.Friends)-1]

			break
		}
	}
}
