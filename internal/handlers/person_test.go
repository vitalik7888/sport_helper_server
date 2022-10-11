package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sport_helper/internal/entity"
	"sport_helper/internal/repository"
	mock_repository "sport_helper/internal/repository/mocks"
	"sport_helper/internal/service"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	dt := []struct {
		name     string
		id       string
		code     int
		err      error
		mockData *entity.Person
	}{
		{"invalid id", "a", http.StatusBadRequest, newInvalidID(nil), nil},
		{"id == 0", "0", http.StatusNotFound, service.ErrInvalidID, nil},
		{"valid id", "1", http.StatusOK, nil,
			&entity.Person{
				ID:        1,
				FirstName: "Fn",
				LastName:  "Ln",
				BirthDate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
				Gender:    "M",
				Height:    175,
			},
		},
		{"No data in repository", "2", http.StatusNotFound, repository.ErrNoContent, nil},
	}

	for _, tc := range dt {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "", nil)
			rec := httptest.NewRecorder()

			id, _ := strconv.Atoi(tc.id)
			repoMock := mock_repository.NewMockPersonRepository(gomock.NewController(t))
			repoMock.EXPECT().GetOne(context.Background(), id).Return(tc.mockData, tc.err).AnyTimes()
			h := NewPersonHandler(service.NewPersonService(repoMock))

			handlerParams := httprouter.Params{httprouter.Param{Key: "id", Value: tc.id}}
			h.(*handlerPerson).getById(rec, req, handlerParams)

			assert.Equal(t, tc.code, rec.Code)
			var data any
			if tc.err != nil {
				data = tc.err
			} else {
				data = tc.mockData
			}
			res, _ := json.Marshal(data)
			se := string(res)
			sa := rec.Body.String()
			assert.Equal(t, se, strings.Replace(sa, "\n", "", -1))
		})
	}
}

func TestGetAll(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	rec := httptest.NewRecorder()

	data := []entity.Person{
		{ID: 1},
		{ID: 2},
	}
	var repo repository.PersonRepository
	repoMock := mock_repository.NewMockPersonRepository(gomock.NewController(t))
	repoMock.EXPECT().GetAll(context.Background()).Return(data, nil)
	repo = repoMock
	h := NewPersonHandler(service.NewPersonService(repo))

	h.(*handlerPerson).getAll(rec, req, nil)

	assert.Equal(t, http.StatusOK, rec.Code)
	res := []entity.Person{}
	b, _ := ioutil.ReadAll(rec.Body)
	json.Unmarshal(b, &res)
	assert.Equal(t, data, res)
}

func TestRemove(t *testing.T) {
	dt := []struct {
		name string
		id   string
		code int
		err  error
	}{
		{"invalid id", "a", http.StatusBadRequest, newInvalidID(nil)},
		{"id == 0", "0", http.StatusBadRequest, service.ErrInvalidID},
		{"valid id", "1", http.StatusOK, nil},
		{"No data in repository", "2", http.StatusOK, nil},
	}

	for _, tc := range dt {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "", nil)
			rec := httptest.NewRecorder()

			id, _ := strconv.Atoi(tc.id)
			repoMock := mock_repository.NewMockPersonRepository(gomock.NewController(t))
			repoMock.EXPECT().Remove(context.Background(), id).Return(tc.err).AnyTimes()
			h := NewPersonHandler(service.NewPersonService(repoMock))

			handlerParams := httprouter.Params{httprouter.Param{Key: "id", Value: tc.id}}
			h.(*handlerPerson).remove(rec, req, handlerParams)

			assert.Equal(t, tc.code, rec.Code)
			var data any
			if tc.err != nil {
				data = tc.err
			} else {
				data = StatusOk
			}
			res, _ := json.Marshal(data)
			se := string(res)
			assert.Equal(t, se, strings.Replace(rec.Body.String(), "\n", "", -1))
		})
	}
}
