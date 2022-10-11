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
	trainingExercisesURL = "/training_exercises"
	trainingExerciseURL  = "/training_exercises/:id"
)

type handlerTrainingExercise struct {
	logger  logger.Logger
	service service.TrainingExerciseService
}

func NewTrainingExerciseHandler(ps service.TrainingExerciseService) Handler {
	return &handlerTrainingExercise{logger: logger.GetLogger(), service: ps}
}

func (h *handlerTrainingExercise) Register(r *httprouter.Router) {
	r.GET(trainingExerciseURL, h.getById)
	r.GET(trainingExercisesURL, h.getAll)
	r.POST(trainingExercisesURL, h.create)
	r.PUT(trainingExerciseURL, h.update)
	r.DELETE(trainingExerciseURL, h.remove)
}

func (h *handlerTrainingExercise) getById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (h *handlerTrainingExercise) getAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	p, err := h.service.GetAll(r.Context())
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (h *handlerTrainingExercise) create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	p := &entity.CreateTrainingExercise{}
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
	json.NewEncoder(w).Encode(&entity.TrainingExercise{
		ID:         lastInsertedId,
		SessionID:  p.SessionID,
		ExerciseID: p.ExerciseID,
		Total:      p.Total,
		Notes:      p.Notes,
	})
}

func (h *handlerTrainingExercise) remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (h *handlerTrainingExercise) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, newInvalidID(err))
		return
	}
	p := &entity.UpdateTrainingExercise{}
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
