package server

import (
	"net/http"
	"sb_social_network/internal/store/teststore"
)

func Start() error {
	db := teststore.New()
	s := NewServer(db)

	return http.ListenAndServe(":8080", s)
}
