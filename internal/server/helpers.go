package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s *server) respondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) checkErr(w http.ResponseWriter, r *http.Request, code int, err error) {
	if err != nil {
		s.respondError(w, r, http.StatusBadRequest, err)
	}
}
