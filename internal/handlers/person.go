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
	personsURL = "/persons"
	personURL  = "/persons/:id"
)

type handlerPerson struct {
	logger  logger.Logger
	service service.PersonService
}

func NewPersonHandler(ps service.PersonService) Handler {
	return &handlerPerson{logger: logger.GetLogger(), service: ps}
}

func (h *handlerPerson) Register(r *httprouter.Router) {
	r.GET(personURL, h.getById)
	r.GET(personsURL, h.getAll)
	r.POST(personsURL, h.create)
	r.PUT(personURL, h.update)
	r.DELETE(personURL, h.remove)
}

func (h *handlerPerson) getById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (h *handlerPerson) getAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	p, err := h.service.GetAll(r.Context())
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (h *handlerPerson) create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	p := &entity.CreatePerson{}
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
	json.NewEncoder(w).Encode(&entity.Person{
		ID:        lastInsertedId,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		BirthDate: p.BirthDate,
		Gender:    p.Gender,
		Height:    p.Height,
	})
}

func (h *handlerPerson) remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (h *handlerPerson) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, newInvalidID(err))
		return
	}
	p := &entity.UpdatePerson{}
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
