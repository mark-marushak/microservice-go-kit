package account

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repository", "sql"),
	}
}

func (r repo) CreateUser(ctx context.Context, user User) error {
	sql := `
	insert into users (id, email, password)
	values ($1, $2, $3)`

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	_, err := r.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := r.db.QueryRow("select * from users where id = $1", id).Scan(email)
	if err != nil {
		return email, err
	}

	return email, nil
}
