package postgres

import (
	"context"
	"database/sql"
	"errors"
	"event-calendar/config"
	"event-calendar/dmodel"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const driverName = "postgres"

var ErrDuplicate = fmt.Errorf("db: duplicate error")
var ErrNoRows = fmt.Errorf("db: no rows in result set")

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepo(db *sql.DB, dbProvider string, DBConfig config.DBConfig) UserRepository {
	//host string, port uint64, user, password, dbname string) (db DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBConfig.Host, DBConfig.Port, DBConfig.Username, DBConfig.Password, DBConfig.DatabaseName)
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

	//type counter struct {
	//	Count int64 `db:"count"`
	//}
	//var count counter
	err := repo.db.GetContext(ctx, &count,
		`SELECT count(*) FROM users`)

	//WHERE nickname ILIKE $1`, fmt.Sprintf("%[1]s%[2]s%[1]s", "%", nickname))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNoRows
		}
		return 0, err
	}

	return count, nil
}

// func (repo UserRepository) GetUsersByID(ctx context.Context, id int64) (users []dmodel.User, err error) {
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

//func (repo UserRepository) FindUserByNickname(nickname string) (user model.User, err error) {
//	err = repo.db.GetFirst(&user, "SELECT * FROM users WHERE nickname = $1", nickname)
//	return user, err
//}
//
//func (repo UserRepository) FindUserByID(id uint64) (user model.User, err error) {
//	err = repo.db.GetFirst(&user, "SELECT * FROM users WHERE id = $1", id)
//	return user, err
//}
//
//func (repo UserRepository) UpdateUserByID(id uint64, newNickname string) error {
//	_, err := repo.db.Connection().Exec("UPDATE users SET nickname = $1, updated_at = now() WHERE id = $2", newNickname, id)
//	return err
//}
