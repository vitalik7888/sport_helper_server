package entity

import (
	"sport_helper/internal/apperror"
	"time"
)

var (
	ErrPersonInvalidID       = apperror.NewEntityError(nil, "invalid person ID")
	ErrPersonInvalidLastName = apperror.NewEntityError(nil, "invalid last name")
)

type CreatePerson struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name,omitempty"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Height    int16     `json:"height"`
}

func (p CreatePerson) Validate() error {
	return validateLastName(p.LastName)
}

type UpdatePerson struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name,omitempty"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Height    int16     `json:"height"`
}

func (p UpdatePerson) Validate() error {
	return validateLastName(p.LastName)
}

type Person struct {
	ID        int       `json:"id,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name,omitempty"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Height    int16     `json:"height"`
}

func (p Person) Validate() error {
	if !validateID(p.ID) {
		return ErrPersonInvalidID
	}

	return validateLastName(p.LastName)
}

func validateID(id int) bool { // FIXME move to separate file
	return id > 0
}

func validateLastName(lastName string) error {
	if lastName == "" {
		return ErrPersonInvalidLastName
	}

	return nil
}

func validateDate(d time.Time) bool {
	return true // FIXME
}