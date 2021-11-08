package repository

import (
	"context"
	"database/sql"
	"github.com/iwdc-oss/monim/server/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindByID(ctx context.Context, tx *sql.Tx, userID int) domain.User
	FindByUsernameAndPassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
