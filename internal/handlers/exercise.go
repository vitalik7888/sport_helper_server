package handlers

import (
	"encoding/json"
	"net/http"
	"sport_helper/internal/entity"
	"sport_helper/internal/service"
	"sport_helper/pkg/logger"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	exercisesURL = "/exercises"
	exerciseURL  = "/exercises/:id"
)

type handlerExercise struct {
	logger  logger.Logger
	service service.ExerciseService
}

func NewExerciseHandler(s service.ExerciseService) Handler {
	return &handlerExercise{logger: logger.GetLogger(), service: s}
}

func (h *handlerExercise) Register(r *httprouter.Router) {
	r.GET(exerciseURL, h.getById)
	r.GET(exercisesURL, h.getAll)
	r.POST(exercisesURL, h.create)
	r.PUT(exerciseURL, h.update)
	r.DELETE(exerciseURL, h.remove)
}

func (h *handlerExercise) getById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, newInvalidID(err))
		return
	}
	p, err := h.service.GetById(r.Context(), id)
	if err != nil {
		newErrorResponse(w, http.StatusNotFound, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (h *handlerExercise) getAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	p, err := h.service.GetAll(r.Context())
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (h *handlerExercise) create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	p := &entity.CreateExercise{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, newInternalError(err))
		return
	}
	lastInsertedId, err := h.service.Create(r.Context(), *p)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&entity.Exercise{
		ID:          lastInsertedId,
		Name:        p.Name,
		Description: p.Description,
	})
}

func (h *handlerExercise) remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, newInvalidID(err))
		return
	}
	if err := h.service.Remove(r.Context(), id); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(StatusOk)
}

func (h *handlerExercise) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, newInvalidID(err))
		return
	}
	p := &entity.UpdateExercise{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, newInternalError(err))
		return
	}
	if err := h.service.Update(r.Context(), id, *p); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(StatusOk)
}
