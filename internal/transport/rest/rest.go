package rest

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Error struct {
	Error string
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

func ReplySuccess(w http.ResponseWriter, status int, payload interface{}) {
	w.WriteHeader(status)

	writeContentType(w, jsonContentType)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Error().Err(err).Msg("json.Encode")
	}
}

func ReplyError(w http.ResponseWriter, status int, err error) {
	log.Error().Err(err).Send()

	w.WriteHeader(status)

	writeContentType(w, jsonContentType)

	if err := json.NewEncoder(w).Encode(Error{
		Error: err.Error(),
	}); err != nil {
		log.Error().Err(err).Msg("json.Encode")
	}
}
