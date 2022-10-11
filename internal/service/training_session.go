package service

import (
	"context"
	"sport_helper/internal/entity"
	"sport_helper/internal/repository"
	"sport_helper/pkg/logger"
)

type TrainingSessionService interface {
	GetById(ctx context.Context, id int) (*entity.TrainingSession, error)
	GetAll(ctx context.Context) ([]entity.TrainingSession, error)
	Create(ctx context.Context, p entity.CreateTrainingSession) (int, error)
	Remove(ctx context.Context, id int) error
	Update(ctx context.Context, id int, p entity.UpdateTrainingSession) error
}

type trainingSessionService struct {
	logger logger.Logger
	repo   repository.TrainingSessionRepository
}

func NewTrainingSessionService(s repository.TrainingSessionRepository) TrainingSessionService {
	return &trainingSessionService{logger: logger.GetLogger(), repo: s}
}

func (s *trainingSessionService) GetById(ctx context.Context, id int) (*entity.TrainingSession, error) {
	if err := validateID(id); err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return s.repo.GetOne(ctx, id)
}

func (s *trainingSessionService) GetAll(ctx context.Context) ([]entity.TrainingSession, error) {
	return s.repo.GetAll(ctx)
}

func (s *trainingSessionService) Create(ctx context.Context, p entity.CreateTrainingSession) (int, error) {
	if err := p.Validate(); err != nil {
		s.logger.Error(err)
		return 0, err
	}
	return s.repo.Create(ctx, p)
}

func (s *trainingSessionService) Remove(ctx context.Context, id int) error {
	if err := validateID(id); err != nil {
		s.logger.Error(err)
		return err
	}
	return s.repo.Remove(ctx, id)
}

func (s *trainingSessionService) Update(ctx context.Context, id int, p entity.UpdateTrainingSession) error {
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
