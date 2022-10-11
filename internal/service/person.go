package service

import (
	"context"
	"sport_helper/internal/apperror"
	"sport_helper/internal/entity"
	"sport_helper/internal/repository"
	"sport_helper/pkg/logger"
)

var (
	ErrInvalidID = apperror.NewServiceError(nil, "invalid id")
)

type PersonService interface {
	GetById(ctx context.Context, id int) (*entity.Person, error)
	GetAll(ctx context.Context) ([]entity.Person, error)
	Create(ctx context.Context, p entity.CreatePerson) (int, error)
	Remove(ctx context.Context, id int) error
	Update(ctx context.Context, id int, p entity.UpdatePerson) error
}

type personService struct {
	logger logger.Logger
	repo   repository.PersonRepository
}

func NewPersonService(s repository.PersonRepository) PersonService {
	return &personService{logger: logger.GetLogger(), repo: s}
}

func (s *personService) GetById(ctx context.Context, id int) (*entity.Person, error) {
	if err := validateID(id); err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return s.repo.GetOne(ctx, id)
}

func (s *personService) GetAll(ctx context.Context) ([]entity.Person, error) {
	return s.repo.GetAll(ctx)
}

func (s *personService) Create(ctx context.Context, p entity.CreatePerson) (int, error) {
	if err := p.Validate(); err != nil {
		s.logger.Error(err)
		return 0, err
	}
	return s.repo.Create(ctx, p)
}

func (s *personService) Remove(ctx context.Context, id int) error {
	if err := validateID(id); err != nil {
		s.logger.Error(err)
		return err
	}
	return s.repo.Remove(ctx, id)
}

func (s *personService) Update(ctx context.Context, id int, p entity.UpdatePerson) error {
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

func validateID(id int) error { // FIXME move to separate file
	if id <= 0 {
		return ErrInvalidID
	}
	return nil
}
