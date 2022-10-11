package entity

import (
	"sport_helper/internal/apperror"
	"time"
)

var (
	ErrTrainSessionInvalidID       = apperror.NewEntityError(nil, "invalid training session ID")
	ErrTrainSessionInvalidPersonID = apperror.NewEntityError(nil, "invalid person id in training session")
	ErrTrainSessionInvalidDate     = apperror.NewEntityError(nil, "invalid training session date")
)

type CreateTrainingSession struct {
	PersonID   int       `json:"person_id,omitempty"`
	Start      time.Time `json:"start,omitempty"`
	End        time.Time `json:"end,omitempty"`
	Evaluation int8      `json:"evaluation"`
	Notes      string    `json:"notes"`
}

func (e CreateTrainingSession) TrainingDuration() time.Duration {
	return duration(e.Start, e.End)
}

func (e CreateTrainingSession) Validate() error {
	return validateStartEndDates(e.Start, e.End)
}

type UpdateTrainingSession struct {
	Start      time.Time `json:"start,omitempty"`
	End        time.Time `json:"end,omitempty"`
	Evaluation int8      `json:"evaluation"`
	Notes      string    `json:"notes"`
}

func (e UpdateTrainingSession) TrainingDuration() time.Duration {
	return duration(e.Start, e.End)
}

func (e UpdateTrainingSession) Validate() error {
	return validateStartEndDates(e.Start, e.End)
}

type TrainingSession struct {
	ID         int       `json:"id,omitempty"`
	PersonID   int       `json:"person_id,omitempty"`
	Start      time.Time `json:"start,omitempty"`
	End        time.Time `json:"end,omitempty"`
	Evaluation int8      `json:"evaluation"`
	Notes      string    `json:"notes"`
}

func (e TrainingSession) TrainingDuration() time.Duration {
	return duration(e.Start, e.End)
}

func (e TrainingSession) Validate() error {
	if !validateID(e.ID) {
		return ErrTrainSessionInvalidID
	}
	if !validateID(e.PersonID) {
		return ErrTrainSessionInvalidPersonID
	}

	return validateStartEndDates(e.Start, e.End)
}

func validateStartEndDates(start time.Time, end time.Time) error {
	// start and end can be equal
	if !validateDate(start) || !validateDate(end) ||
		start.Before(end) {
		return ErrTrainSessionInvalidDate
	}
	return nil
}

func duration(start time.Time, end time.Time) time.Duration {
	return end.Sub(start)
}
