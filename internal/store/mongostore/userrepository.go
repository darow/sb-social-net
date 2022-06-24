package mongostore

import (
	"context"
	"errors"
	"fmt"
	"log"

	"sb_social_network/internal/models"
	"sb_social_network/internal/store"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	store *Store
	users *mongo.Collection
}

func (r *UserRepository) Create(user models.User) string {
	res, err := r.users.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	oid, _ := res.InsertedID.(primitive.ObjectID)

	return fmt.Sprintf("%s", oid.Hex())
}

func (r *UserRepository) MakeFriends(user1, user2 models.User) {
	if !r.contains(user1.Friends, user2.ID) {
		user1.Friends = append(user1.Friends, user2.ID)
		update := bson.D{
			{"$set", bson.D{{"friends", user1.Friends}}},
		}
		_, err := r.users.UpdateOne(context.TODO(), bson.D{{"_id", user1.ID}}, update)
		if err != nil {
			fmt.Println(err)
			return
		}

		user2.Friends = append(user2.Friends, user1.ID)
		update = bson.D{
			{"$set", bson.D{{"friends", user2.Friends}}},
		}
		_, err = r.users.UpdateOne(context.TODO(), bson.D{{"_id", user2.ID}}, update)
		if err != nil {
			fmt.Println(err)
		}
	}

	return
}

// Delete Удаляем пользователя и ссылки на дружбу с ним у всех пользователей.
func (r *UserRepository) Delete(user *models.User) {
	filter := bson.D{
		{"friends", bson.D{{"$all", []primitive.ObjectID{user.ID}}}},
	}
	cursor, err := r.users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var friends []models.User
	if err = cursor.All(context.TODO(), &friends); err != nil {
		log.Fatal(err)
	}

	for _, f := range friends {
		i := 0
		for ; i < len(f.Friends); i++ {
			if f.Friends[i] == user.ID {
				break
			}
		}

		f.Friends = append(f.Friends[:i], f.Friends[i+1:]...)
		filter = bson.D{{"_id", f.ID}}
		update := bson.D{{"$set", bson.D{{"friends", f.Friends}}}}
		_, err = r.users.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			fmt.Println(err)
		}
	}

	filter = bson.D{{"_id", user.ID}}
	r.users.DeleteOne(context.TODO(), filter)
}

func (r *UserRepository) SetAge(user *models.User, newAge int) {
	filter := bson.D{{"_id", user.ID}}
	update := bson.D{{"$set", bson.D{{"age", newAge}}}}
	r.users.UpdateOne(context.TODO(), filter, update)
}

func (r *UserRepository) contains(s []primitive.ObjectID, e primitive.ObjectID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", oid}}

	u := models.User{}
	err = r.users.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, store.ErrObjectNotFound
		}
		return nil, err
	}

	return &u, nil
}
