package rest

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Error struct {
	Error string
}

var jsonContentType = "application/json; charset=utf-8"

func ReplySuccess(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", jsonContentType)

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Error().Err(err).Msg("json.Encode")
	}
}

func ReplyError(w http.ResponseWriter, status int, err error) {
	log.Error().Err(err).Send()

	w.Header().Set("Content-Type", jsonContentType)

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(Error{
		Error: err.Error(),
	}); err != nil {
		log.Error().Err(err).Msg("json.Encode")
	}
}
