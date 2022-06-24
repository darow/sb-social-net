package server

import (
	"context"
	"fmt"
	"net/http"

	"sb_social_network/internal/store/mongostore"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Start(config *Config) error {
	db, err := newDB(config.DB)
	if err != nil {
		return err
	}

	store := mongostore.New(db)
	s := NewServer(store)

	return http.ListenAndServe(":8080", s)
}

func newDB(cfg ConfigDB) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}
	fmt.Println("connected to mongo successfully")

	db := client.Database(cfg.Name)

	return db, nil
}
