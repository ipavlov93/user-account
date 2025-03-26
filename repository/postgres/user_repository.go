package postgres

import (
	"context"
	"database/sql"
	"errors"
	"event-calendar/dmodel"
	"event-calendar/repository"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const driverName = "postgres"

var ErrDuplicate = fmt.Errorf("db: duplicate error")
var ErrNoRows = fmt.Errorf("db: no rows in result set")

type UserRepository struct {
	db *sqlx.DB
}

var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepo(host string, port int64, user, password, dbname string) UserRepository {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//DBConfig.Host, DBConfig.Port, DBConfig.Username, DBConfig.Password, DBConfig.DatabaseName)
	connection, err := sqlx.Connect(driverName, psqlInfo)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database, %s", err))
	}

	return UserRepository{
		db: connection,
	}
}

func (repo UserRepository) GetUsersCount(ctx context.Context) (int64, error) {
	var count int64
	err := repo.db.GetContext(ctx, &count,
		`SELECT count(*) FROM users`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNoRows
		}
		return 0, err
	}

	return count, nil
}

func (repo UserRepository) GetUserByID(ctx context.Context, id int64) (user dmodel.User, err error) {
	err = repo.db.GetContext(ctx, &user,
		`SELECT * FROM users
				WHERE id = $1`, id)
	return user, err
}

func (repo UserRepository) GetUserByUUID(ctx context.Context, uuid string) (user dmodel.User, err error) {
	err = repo.db.GetContext(ctx, &user,
		`SELECT * FROM users
				WHERE uuid = $1`, uuid)
	return user, err
}

func (repo UserRepository) AddUser(ctx context.Context, user dmodel.User) (id int64, err error) {
	err = repo.db.QueryRowContext(
		ctx,
		`INSERT INTO users (uuid, first_name, last_name, email_address, description)
				VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		user.UUID, user.FirstName, user.LastName, user.EmailAddress, user.Description,
	).Scan(&id)
	if err != nil {
		if len(err.Error()) > 50 {
			if err.Error()[:50] == "pq: duplicate key value violates unique constraint" {
				return 0, ErrDuplicate
			}
		}
		return 0, err
	}
	return id, nil
}
