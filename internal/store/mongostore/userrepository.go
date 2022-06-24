package mongostore

import (
	"context"
	"errors"
	"fmt"

	"sb_social_network/internal/model"
	"sb_social_network/internal/store"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	store *Store
	users *mongo.Collection
}

func (r *UserRepository) Create(user model.User) string {
	res, err := r.users.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	oid, _ := res.InsertedID.(primitive.ObjectID)

	return fmt.Sprintf("%s", oid.Hex())
}

func (r *UserRepository) MakeFriends(user1, user2 model.User) {
	if !r.contains(user1.Friends, user2.ID) {
		user1.Friends = append(user1.Friends, user2.ID)
		oid1, _ := primitive.ObjectIDFromHex(user1.ID)
		update := bson.D{
			{"$set", bson.D{{"friends", user1.Friends}}},
		}
		_, err := r.users.UpdateOne(context.TODO(), bson.D{{"_id", oid1}}, update)
		if err != nil {
			fmt.Println(err)
			return
		}

		user2.Friends = append(user2.Friends, user1.ID)
		oid2, _ := primitive.ObjectIDFromHex(user2.ID)
		update = bson.D{
			{"$set", bson.D{{"friends", user2.Friends}}},
		}
		_, err = r.users.UpdateOne(context.TODO(), bson.D{{"_id", oid2}}, update)
		if err != nil {
			fmt.Println(err)
		}
	}

	return
}

// Delete Удаляем пользователя и ссылки на дружбу с ним у всех пользователей.
func (r *UserRepository) Delete(user *model.User) {
	//for _, u := range r.users {
	//	u.RemoveFromFriends(user)
	//}
	//delete(r.users, user.ID)
}

func (r *UserRepository) SetAge(user *model.User, newAge int) error {
	//u, ok := r.users[user.ID]
	//if !ok {
	//	return ErrObjectNotFound
	//}
	//u.Age = newAge
	//
	//return nil
	return errors.New("")
}

func (r *UserRepository) contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func (r *UserRepository) FindByID(id string) (*model.User, error) {
	oid := primitive.ObjectID{}
	oid.UnmarshalText([]byte(id))
	filter := bson.D{{"_id", oid}}
	u := model.User{}
	err := r.users.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, store.ErrObjectNotFound
		}
		return nil, err
	}
	u.ID = id

	return &u, nil
}
