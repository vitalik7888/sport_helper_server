package handlers

import (
	"encoding/json"
	"net/http"
	"sport_helper/internal/apperror"
	"sport_helper/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	Register(*httprouter.Router)
}

var (
	StatusOk = &status{"ok"}
)

type status struct {
	Message string `json:"status,omitempty"`
}

func newInvalidID(err error) *apperror.AppError {
	return apperror.NewHandlerError(err, "invalid id")
}

func newInternalError(err error) *apperror.AppError {
	return apperror.NewHandlerError(err, "system error")
}

func newErrorResponse(w http.ResponseWriter, code int, err error) {
	logger := logger.GetLogger()
	logger.Error(err)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
