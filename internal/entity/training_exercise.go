package entity

import "sport_helper/internal/apperror"

var (
	ErrTrainExerInvalidID         = apperror.NewEntityError(nil, "invalid training exercise ID")
	ErrTrainExerInvalidSessionID  = apperror.NewEntityError(nil, "invalid session id in training exercise")
	ErrTrainExerInvalidExerciseID = apperror.NewEntityError(nil, "invalid exercise id in training exercise")
	ErrTrainExerTotal             = apperror.NewEntityError(nil, "invalid training exercise total value")
)

type CreateTrainingExercise struct {
	SessionID  int    `json:"session_id,omitempty"`
	ExerciseID int    `json:"exercise_id,omitempty"`
	Total      string `json:"total,omitempty"`
	Notes      string `json:"notes"`
}

func (e CreateTrainingExercise) Validate() error {
	if !validateID(e.SessionID) {
		return ErrTrainExerInvalidSessionID
	}
	if !validateID(e.ExerciseID) {
		return ErrTrainExerInvalidExerciseID
	}
	return validateTotal(e.Total)
}

type UpdateTrainingExercise struct {
	SessionID  int    `json:"session_id,omitempty"`
	ExerciseID int    `json:"exercise_id,omitempty"`
	Total      string `json:"total,omitempty"`
	Notes      string `json:"notes"`
}

func (e UpdateTrainingExercise) Validate() error {
	if !validateID(e.ExerciseID) {
		return ErrTrainExerInvalidSessionID
	}
	if !validateID(e.SessionID) {
		return ErrTrainExerInvalidSessionID
	}
	return validateTotal(e.Total)
}

type TrainingExercise struct {
	ID         int    `json:"id,omitempty"`
	SessionID  int    `json:"session_id,omitempty"`
	ExerciseID int    `json:"exercise_id,omitempty"`
	Total      string `json:"total,omitempty"`
	Notes      string `json:"notes"`
}

func (e TrainingExercise) Validate() error {
	if !validateID(e.ID) {
		return ErrTrainExerInvalidID
	}
	if !validateID(e.SessionID) {
		return ErrTrainExerInvalidSessionID
	}
	if !validateID(e.ExerciseID) {
		return ErrTrainExerInvalidExerciseID
	}
	return validateTotal(e.Total)
}

func validateTotal(value string) error {
	if value == "" {
		return ErrTrainExerTotal
	}
	return nil
}
