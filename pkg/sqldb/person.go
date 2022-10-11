package sqldb

import (
	"context"
	"database/sql"
	"sport_helper/internal/apperror"
	"sport_helper/internal/entity"
	"sport_helper/internal/repository"
	"sport_helper/pkg/logger"
)

type personRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func (pr *personRepository) GetOne(ctx context.Context, id int) (*entity.Person, error) {
	row := pr.db.QueryRowContext(ctx, "SELECT * FROM persons WHERE id=?", id)

	var p entity.Person
	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.BirthDate, &p.Gender, &p.Height)
	if err != nil {
		pr.logger.Error(err)
		return nil, repository.ErrNoContent
	}

	return &p, nil
}

func (pr *personRepository) GetAll(ctx context.Context) ([]entity.Person, error) {
	rows, err := pr.db.QueryContext(ctx, "SELECT * FROM persons")
	if err != nil {
		pr.logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	data := []entity.Person{}
	for rows.Next() {
		p := entity.Person{}
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.BirthDate, &p.Gender, &p.Height)
		if err != nil {
			pr.logger.Error(err)
			return nil, apperror.NewRepoError(err, "can`t get data")
		}
		data = append(data, p)
	}
	return data, nil
}

func (pr *personRepository) Create(ctx context.Context, p entity.CreatePerson) (int, error) {
	res, err := pr.db.ExecContext(ctx,
		"INSERT INTO persons(first_name, last_name, birth_date, gender, height) VALUES (?, ?, ?, ?, ?);",
		p.FirstName, p.LastName, p.BirthDate, p.Gender, p.Height,
	)
	if err != nil {
		pr.logger.Error(err)
		return 0, apperror.NewRepoError(err, "can`t create person")
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		pr.logger.Error(err)
		return 0, apperror.NewRepoError(err, "can`t get last created person id")
	}
	return int(id), nil
}

func (pr *personRepository) Remove(ctx context.Context, id int) error {
	_, err := pr.db.ExecContext(ctx, "DELETE FROM persons WHERE id=?", id)
	if err != nil {
		pr.logger.Error(err)
		return apperror.NewRepoError(err, "can`t remove person")
	}
	return nil
}

func (pr *personRepository) Update(ctx context.Context, id int, p entity.UpdatePerson) error {
	_, err := pr.db.ExecContext(ctx,
		"UPDATE persons SET first_name = ?, last_name = ?, birth_date = ?, gender = ?, height = ? WHERE id=?;",
		p.FirstName, p.LastName, p.BirthDate, p.Gender, p.Height, id)
	if err != nil {
		pr.logger.Error(err)
		return apperror.NewRepoError(err, "can`t update person")
	}
	return nil
}