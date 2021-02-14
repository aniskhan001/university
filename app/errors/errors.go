package errors

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	ErrNotFound            = errors.New("resource not found")
	ErrConflict            = errors.New("data conflict or already exist")
	ErrBadRequest          = errors.New("bad request, check param or body")
	ErrInternalServerError = errors.New("internal server error")
)

func getStatusCode(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrInternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

type ErrResponse struct {
	Error string `json:"error"`
}

// RespondError takes an `error` and a `customErr message` args
// to log the error to system and return to client
func RespondError(err error, customErr ...error) (int, ErrResponse) {
	logrus.Errorln(err, customErr)
	if len(customErr) != 0 {
		return getStatusCode(err), ErrResponse{Error: customErr[0].Error()}
	}
	return getStatusCode(err), ErrResponse{Error: err.Error()}
}
