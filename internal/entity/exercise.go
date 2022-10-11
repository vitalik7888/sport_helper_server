package entity

import (
	"sport_helper/internal/apperror"
)

var (
	ErrExerciseInvalidID   = apperror.NewEntityError(nil, "invalid exercise ID")
	ErrExerciseInvalidName = apperror.NewEntityError(nil, "invalid exercise name")
)

type CreateExercise struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
}

func (p CreateExercise) Validate() error {
	return validateName(p.Name)
}

type UpdateExercise struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
}

func (p UpdateExercise) Validate() error {
	return validateName(p.Name)
}

type Exercise struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
}

func (p Exercise) Validate() error {
	if !validateID(p.ID) {
		return ErrExerciseInvalidID
	}

	return validateName(p.Name)
}

func validateName(name string) error {
	if name == "" {
		return ErrExerciseInvalidName
	}

	return nil
}
