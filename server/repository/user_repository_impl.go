package repository

import (
	"context"
	"database/sql"

	"github.com/iwdc-oss/monim/server/helper"
	"github.com/iwdc-oss/monim/server/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO user(username, email, password, first_name, last_name) VALUES(?, ?, ?, ?, ?)"
	res, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password, user.FirstName, user.LastName)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	user.ID = int32(id)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE user SET username = ?, email = ?, first_name = ?, last_name = ?, profile_image = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.FirstName, user.LastName, user.ProfileImage, user.ID)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, userID int) domain.User {
	SQL := "SELECT id, username, email, first_name, last_name, profile_image FROM user WHERE id = ? LIMIT 1"

	rows, err := tx.QueryContext(ctx, SQL, userID)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.User
	err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.ProfileImage)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) FindByUsernameAndPassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "SELECT id, username, email, first_name, last_name, profile_image FROM user WHERE (username = ? AND password = ?) LIMIT 1"

	rows, err := tx.QueryContext(ctx, SQL, user.Username, user.Password)
	helper.PanicIfError(err)
	defer rows.Close()

	err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.ProfileImage)
	helper.PanicIfError(err)

	return user
}
