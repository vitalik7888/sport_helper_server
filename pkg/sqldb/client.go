package sqldb

import (
	"database/sql"
	"sport_helper/internal/repository"
	"sport_helper/pkg/logger"

	_ "github.com/mattn/go-sqlite3"
)

type DbClient struct {
	logger               logger.Logger
	db                   *sql.DB
	personRepo           *personRepository
	exerciseRepo         *exerciseRepository
	trainingSessionRepo  *trainingSessionRepository
	trainingExerciseRepo *trainingExerciseRepository
}

func NewSqliteClient(constr string) (*DbClient, error) {
	logger := logger.GetLogger()
	logger.Info("Creating sqlite client")

	db, dbErr := sql.Open("sqlite3", constr)
	if dbErr != nil {
		return nil, dbErr
	}
	client := &DbClient{
		logger:               logger,
		db:                   db,
		personRepo:           &personRepository{logger: logger, db: db},
		exerciseRepo:         &exerciseRepository{logger: logger, db: db},
		trainingSessionRepo:  &trainingSessionRepository{logger: logger, db: db},
		trainingExerciseRepo: &trainingExerciseRepository{logger: logger, db: db},
	}
	return client, nil
}

func (c *DbClient) Persons() repository.PersonRepository {
	return c.personRepo
}

func (c *DbClient) Exercise() repository.ExerciseRepository {
	return c.exerciseRepo
}

func (c *DbClient) TrainingSession() repository.TrainingSessionRepository {
	return c.trainingSessionRepo
}

func (c *DbClient) TrainingExercise() repository.TrainingExerciseRepository {
	return c.trainingExerciseRepo
}
