package service

import (
	"context"
	"sport_helper/internal/entity"
	"sport_helper/internal/repository"
	"sport_helper/pkg/logger"
)

type TrainingExerciseService interface {
	GetById(ctx context.Context, id int) (*entity.TrainingExercise, error)
	GetAll(ctx context.Context) ([]entity.TrainingExercise, error)
	Create(ctx context.Context, p entity.CreateTrainingExercise) (int, error)
	Remove(ctx context.Context, id int) error
	Update(ctx context.Context, id int, p entity.UpdateTrainingExercise) error
}

type trainingExerciseService struct {
	logger logger.Logger
	repo   repository.TrainingExerciseRepository
}

func NewTrainingExerciseService(s repository.TrainingExerciseRepository) TrainingExerciseService {
	return &trainingExerciseService{logger: logger.GetLogger(), repo: s}
}

func (s *trainingExerciseService) GetById(ctx context.Context, id int) (*entity.TrainingExercise, error) {
	if err := validateID(id); err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return s.repo.GetOne(ctx, id)
}

func (s *trainingExerciseService) GetAll(ctx context.Context) ([]entity.TrainingExercise, error) {
	return s.repo.GetAll(ctx)
}

func (s *trainingExerciseService) Create(ctx context.Context, p entity.CreateTrainingExercise) (int, error) {
	if err := p.Validate(); err != nil {
		s.logger.Error(err)
		return 0, err
	}
	return s.repo.Create(ctx, p)
}

func (s *trainingExerciseService) Remove(ctx context.Context, id int) error {
	if err := validateID(id); err != nil {
		s.logger.Error(err)
		return err
	}
	return s.repo.Remove(ctx, id)
}

func (s *trainingExerciseService) Update(ctx context.Context, id int, p entity.UpdateTrainingExercise) error {
	if err := validateID(id); err != nil {
		s.logger.Error(err)
		return err
	}
	if err := p.Validate(); err != nil {
		s.logger.Error(err)
		return err
	}
	return s.repo.Update(ctx, id, p)
}
