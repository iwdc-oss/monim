package repository

import (
	"context"
	"database/sql"
	"github.com/iwdc-oss/monim/server/model/domain"
)

type SosmedRepository interface {
	Create(ctx context.Context, tx *sql.Tx, sosmed domain.Sosmed) domain.Sosmed
	Update(ctx context.Context, tx *sql.Tx, sosmed domain.Sosmed) domain.Sosmed
	FindByID(ctx context.Context, tx *sql.Tx, sosmedID int) domain.Sosmed
}
