package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty"`
	Name    string               `bson:"name"`
	Age     int                  `bson:"age"`
	Friends []primitive.ObjectID `bson:"friends"`
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
