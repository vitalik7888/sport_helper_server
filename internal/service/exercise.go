package service

import (
	"context"
	"sport_helper/internal/entity"
	"sport_helper/internal/repository"
	"sport_helper/pkg/logger"
)

type ExerciseService interface {
	GetById(ctx context.Context, id int) (*entity.Exercise, error)
	GetAll(ctx context.Context) ([]entity.Exercise, error)
	Create(ctx context.Context, p entity.CreateExercise) (int, error)
	Remove(ctx context.Context, id int) error
	Update(ctx context.Context, id int, p entity.UpdateExercise) error
}

type exerciseService struct {
	logger logger.Logger
	repo   repository.ExerciseRepository
}

func NewExerciseService(s repository.ExerciseRepository) ExerciseService {
	return &exerciseService{logger: logger.GetLogger(), repo: s}
}

func (s *exerciseService) GetById(ctx context.Context, id int) (*entity.Exercise, error) {
	if err := validateID(id); err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return s.repo.GetOne(ctx, id)
}

func (s *exerciseService) GetAll(ctx context.Context) ([]entity.Exercise, error) {
	return s.repo.GetAll(ctx)
}

func (s *exerciseService) Create(ctx context.Context, p entity.CreateExercise) (int, error) {
	if err := p.Validate(); err != nil {
		s.logger.Error(err)
		return 0, err
	}
	return s.repo.Create(ctx, p)
}

func (s *exerciseService) Remove(ctx context.Context, id int) error {
	if err := validateID(id); err != nil {
		s.logger.Error(err)
		return err
	}
	return s.repo.Remove(ctx, id)
}

func (s *exerciseService) Update(ctx context.Context, id int, p entity.UpdateExercise) error {
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
