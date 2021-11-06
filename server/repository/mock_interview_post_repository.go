package repository

import (
	"context"
	"database/sql"

	"github.com/iwdc-oss/monim/server/model/domain"
)

type MockInterviewPostRepository interface {
	Create(ctx context.Context, tx *sql.Tx, mockInterviewPost domain.MockInterviewPost) domain.MockInterviewPost
	FindById(ctx context.Context, tx *sql.Tx, id int32) domain.MockInterviewPost
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MockInterviewPost
}
