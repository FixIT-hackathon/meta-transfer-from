package responses

import (
	"net/http"
	"github.com/google/jsonapi"
)

type Error struct {
	ErrorType string `json:"error_type"`
	ErrorCode int    `json:"error_code"`
	ErrorBody string `json:"error_body"`
}

func WriteError(w http.ResponseWriter, code int, objects []*jsonapi.ErrorObject) {
	w.WriteHeader(code)
	jsonapi.MarshalErrors(w, objects)
}
