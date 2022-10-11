package sqldb

import (
	"context"
	"database/sql"
	"sport_helper/internal/apperror"
	"sport_helper/internal/entity"
	"sport_helper/internal/repository"
	"sport_helper/pkg/logger"
)

type exerciseRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func (pr *exerciseRepository) GetOne(ctx context.Context, id int) (*entity.Exercise, error) {
	row := pr.db.QueryRowContext(ctx, "SELECT * FROM exercises WHERE id=?", id)

	var p entity.Exercise
	err := row.Scan(&p.ID, &p.Name, &p.Description)
	if err != nil {
		pr.logger.Error(err)
		return nil, repository.ErrNoContent
	}

	return &p, nil
}

func (pr *exerciseRepository) GetAll(ctx context.Context) ([]entity.Exercise, error) {
	rows, err := pr.db.QueryContext(ctx, "SELECT * FROM exercises")
	if err != nil {
		pr.logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	data := []entity.Exercise{}
	for rows.Next() {
		p := entity.Exercise{}
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			pr.logger.Error(err)
			return nil, apperror.NewRepoError(err, "can`t get data")
		}
		data = append(data, p)
	}
	return data, nil
}

func (pr *exerciseRepository) Create(ctx context.Context, e entity.CreateExercise) (int, error) {
	res, err := pr.db.ExecContext(ctx,
		"INSERT INTO exercises(name, description) VALUES (?, ?);",
		e.Name, e.Description,
	)
	if err != nil {
		pr.logger.Error(err)
		return 0, apperror.NewRepoError(err, "can`t create exercise")
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		pr.logger.Error(err)
		return 0, apperror.NewRepoError(err, "can`t get last created exercise id")
	}
	return int(id), nil
}

func (pr *exerciseRepository) Remove(ctx context.Context, id int) error {
	_, err := pr.db.ExecContext(ctx, "DELETE FROM exercises WHERE id=?", id)
	if err != nil {
		pr.logger.Error(err)
		return apperror.NewRepoError(err, "can`t remove exercise")
	}
	return nil
}

func (pr *exerciseRepository) Update(ctx context.Context, id int, e entity.UpdateExercise) error {
	_, err := pr.db.ExecContext(ctx,
		"UPDATE exercises SET name = ?, description = ? WHERE id=?;",
		e.Name, e.Description, id)
	if err != nil {
		pr.logger.Error(err)
		return apperror.NewRepoError(err, "can`t update exercise")
	}
	return nil
}