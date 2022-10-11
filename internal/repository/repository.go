package repository

import (
	"context"
	"sport_helper/internal/apperror"
	"sport_helper/internal/entity"
)

var (
	ErrNoContent = apperror.NewRepoError(nil, "no content")
)

type PersonRepository interface {
	GetOne(ctx context.Context, id int) (*entity.Person, error)
	GetAll(ctx context.Context) ([]entity.Person, error)
	Create(ctx context.Context, p entity.CreatePerson) (int, error)
	Remove(ctx context.Context, id int) error
	Update(ctx context.Context, id int, p entity.UpdatePerson) error
}

type ExerciseRepository interface {
	GetOne(ctx context.Context, id int) (*entity.Exercise, error)
	GetAll(ctx context.Context) ([]entity.Exercise, error)
	Create(ctx context.Context, p entity.CreateExercise) (int, error)
	Remove(ctx context.Context, id int) error
	Update(ctx context.Context, id int, p entity.UpdateExercise) error
}

type TrainingSessionRepository interface {
	GetOne(ctx context.Context, id int) (*entity.TrainingSession, error)
	GetAll(ctx context.Context) ([]entity.TrainingSession, error)
	Create(ctx context.Context, p entity.CreateTrainingSession) (int, error)
	Remove(ctx context.Context, id int) error
	Update(ctx context.Context, id int, p entity.UpdateTrainingSession) error
}

type TrainingExerciseRepository interface {
	GetOne(ctx context.Context, id int) (*entity.TrainingExercise, error)
	GetAll(ctx context.Context) ([]entity.TrainingExercise, error)
	Create(ctx context.Context, p entity.CreateTrainingExercise) (int, error)
	Remove(ctx context.Context, id int) error
	Update(ctx context.Context, id int, p entity.UpdateTrainingExercise) error
}

type Respositories struct {
}
