package server

import "errors"

var (
	ErrCtxDoesNotExist = errors.New("object with ctx key does not exist")
)
