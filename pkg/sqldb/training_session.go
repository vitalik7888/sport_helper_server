package sqldb

import (
	"context"
	"database/sql"
	"sport_helper/internal/apperror"
	"sport_helper/internal/entity"
	"sport_helper/internal/repository"
	"sport_helper/pkg/logger"
)

type trainingSessionRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func (pr *trainingSessionRepository) GetOne(ctx context.Context, id int) (*entity.TrainingSession, error) {
	row := pr.db.QueryRowContext(ctx, "SELECT * FROM training_sessions WHERE id=?;", id)

	var e entity.TrainingSession
	err := row.Scan(&e.ID, &e.PersonID, &e.Start, &e.End, &e.Evaluation, &e.Notes)
	if err != nil {
		pr.logger.Error(err)
		return nil, repository.ErrNoContent
	}

	return &e, nil
}

func (pr *trainingSessionRepository) GetAll(ctx context.Context) ([]entity.TrainingSession, error) {
	rows, err := pr.db.QueryContext(ctx, "SELECT * FROM training_sessions;")
	if err != nil {
		pr.logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	data := []entity.TrainingSession{}
	for rows.Next() {
		e := entity.TrainingSession{}
		err := rows.Scan(&e.ID, &e.PersonID, &e.Start, &e.End, &e.Evaluation, &e.Notes)
		if err != nil {
			pr.logger.Error(err)
			return nil, apperror.NewRepoError(err, "can`t get data")
		}
		data = append(data, e)
	}
	return data, nil
}

func (pr *trainingSessionRepository) Create(ctx context.Context, e entity.CreateTrainingSession) (int, error) {
	res, err := pr.db.ExecContext(ctx,
		"INSERT INTO training_sessions(person_id, start, end, evaluation, notes) VALUES (?, ?, ?, ?, ?);",
		e.PersonID, e.Start, e.End, e.Evaluation, e.Notes,
	)
	if err != nil {
		pr.logger.Error(err)
		return 0, apperror.NewRepoError(err, "can`t create training session entity")
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		pr.logger.Error(err)
		return 0, apperror.NewRepoError(err, "can`t get last created training session id")
	}
	return int(id), nil
}

func (pr *trainingSessionRepository) Remove(ctx context.Context, id int) error {
	_, err := pr.db.ExecContext(ctx, "DELETE FROM training_sessions WHERE id=?;", id)
	if err != nil {
		pr.logger.Error(err)
		return apperror.NewRepoError(err, "can`t remove training_session")
	}
	return nil
}

func (pr *trainingSessionRepository) Update(ctx context.Context, id int, e entity.UpdateTrainingSession) error {
	_, err := pr.db.ExecContext(ctx,
		"UPDATE training_sessions SET start=?, end=?, evaluation=?, notes=? WHERE id=?;",
		e.Start, e.End, e.Evaluation, e.Notes, id)
	if err != nil {
		pr.logger.Error(err)
		return apperror.NewRepoError(err, "can`t update training_session")
	}
	return nil
}