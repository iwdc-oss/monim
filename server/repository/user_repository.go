package repository

import (
	"context"
	"database/sql"
	"monim/server/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, id int32) domain.User
	FindByUsernameAndPassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
