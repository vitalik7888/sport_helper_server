package sqldb

import (
	"context"
	"database/sql"
	"sport_helper/internal/apperror"
	"sport_helper/internal/entity"
	"sport_helper/internal/repository"
	"sport_helper/pkg/logger"
)

type trainingExerciseRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func (r *trainingExerciseRepository) GetOne(ctx context.Context, id int) (*entity.TrainingExercise, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM training_exercises WHERE id=?", id)

	var e entity.TrainingExercise
	err := row.Scan(&e.ID, &e.SessionID, &e.ExerciseID, &e.Total, &e.Notes)
	if err != nil {
		r.logger.Error(err)
		return nil, repository.ErrNoContent
	}

	return &e, nil
}

func (r *trainingExerciseRepository) GetAll(ctx context.Context) ([]entity.TrainingExercise, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM training_exercises")
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	data := []entity.TrainingExercise{}
	for rows.Next() {
		e := entity.TrainingExercise{}
		err := rows.Scan(&e.ID, &e.SessionID, &e.ExerciseID, &e.Total, &e.Notes)
		if err != nil {
			r.logger.Error(err)
			return nil, apperror.NewRepoError(err, "can`t get data")
		}
		data = append(data, e)
	}
	return data, nil
}

func (r *trainingExerciseRepository) Create(ctx context.Context, e entity.CreateTrainingExercise) (int, error) {
	res, err := r.db.ExecContext(ctx,
		"INSERT INTO training_exercises(session_id, exercise_id, total, notes) VALUES (?, ?, ?, ?);",
		e.SessionID, e.ExerciseID, e.Total, e.Notes,
	)
	if err != nil {
		r.logger.Error(err)
		return 0, apperror.NewRepoError(err, "can`t create training exercise")
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		r.logger.Error(err)
		return 0, apperror.NewRepoError(err, "can`t get last created training exercise id")
	}
	return int(id), nil
}

func (r *trainingExerciseRepository) Remove(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM training_exercises WHERE id=?", id)
	if err != nil {
		r.logger.Error(err)
		return apperror.NewRepoError(err, "can`t remove training exercise")
	}
	return nil
}

func (r *trainingExerciseRepository) Update(ctx context.Context, id int, e entity.UpdateTrainingExercise) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE training_exercises SET session_id=?, exercise_id=?, total=?, notes=? WHERE id=?;",
		e.SessionID, e.ExerciseID, e.Total, e.Notes, id,
	)
	if err != nil {
		r.logger.Error(err)
		return apperror.NewRepoError(err, "can`t update training exercise")
	}
	return nil
}