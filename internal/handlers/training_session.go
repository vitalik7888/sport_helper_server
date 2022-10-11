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
	trainingSessionsURL = "/training_sessions"
	trainingSessionURL  = "/training_sessions/:id"
)

type handlerTrainingSession struct {
	logger  logger.Logger
	service service.TrainingSessionService
}

func NewTrainingSessionHandler(ps service.TrainingSessionService) Handler {
	return &handlerTrainingSession{logger: logger.GetLogger(), service: ps}
}

func (h *handlerTrainingSession) Register(r *httprouter.Router) {
	r.GET(trainingSessionURL, h.getById)
	r.GET(trainingSessionsURL, h.getAll)
	r.POST(trainingSessionsURL, h.create)
	r.PUT(trainingSessionURL, h.update)
	r.DELETE(trainingSessionURL, h.remove)
}

func (h *handlerTrainingSession) getById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (h *handlerTrainingSession) getAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	p, err := h.service.GetAll(r.Context())
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (h *handlerTrainingSession) create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	p := &entity.CreateTrainingSession{}
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
	json.NewEncoder(w).Encode(&entity.TrainingSession{
		ID:         lastInsertedId,
		PersonID:   p.PersonID,
		Start:      p.Start,
		End:        p.End,
		Evaluation: p.Evaluation,
		Notes:      p.Notes,
	})
}

func (h *handlerTrainingSession) remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (h *handlerTrainingSession) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, newInvalidID(err))
		return
	}
	p := &entity.UpdateTrainingSession{}
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
