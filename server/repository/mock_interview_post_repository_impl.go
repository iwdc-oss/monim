package repository

import (
	"context"
	"database/sql"

	"github.com/iwdc-oss/monim/server/helper"
	"github.com/iwdc-oss/monim/server/model/domain"
)

type MockInterviewPostRepositoryImpl struct{}

func NewMockInterviewPostRepositoryImpl() *MockInterviewPostRepositoryImpl {
	return &MockInterviewPostRepositoryImpl{}
}

func (repository *MockInterviewPostRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, mockInterviewPost domain.MockInterviewPost) domain.MockInterviewPost {
	SQL := "INSERT INTO mock_interview_post(name, role, logo, description, agreement) VALUES(?, ?, ?, ?, ?)"
	res, err := tx.ExecContext(ctx, SQL, mockInterviewPost.Name, mockInterviewPost.Role, mockInterviewPost.Logo, mockInterviewPost.Description, mockInterviewPost.Agreement)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	mockInterviewPost.Id = int32(id)

	return mockInterviewPost
}

func (repository *MockInterviewPostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int32) domain.MockInterviewPost {
	SQL := "SELECT name, role, logo, description, agreement, published_at FROM mock_interview_post WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	mockInterviewPost := domain.MockInterviewPost{}
	rows.Scan(&mockInterviewPost.Name, &mockInterviewPost.Role, &mockInterviewPost.Logo, &mockInterviewPost.Description, &mockInterviewPost.Agreement, &mockInterviewPost.PublishedAt)
	mockInterviewPost.Id = id

	return mockInterviewPost

}

func (repository *MockInterviewPostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MockInterviewPost {
	SQL := "SELECT id, name, role, logo, description, agreement, published_at FROM mock_interview_post"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var mockInterviewPostSlice []domain.MockInterviewPost
	for rows.Next() {
		mockInterviewPost := domain.MockInterviewPost{}

		rows.Scan(&mockInterviewPost.Id, &mockInterviewPost.Name, &mockInterviewPost.Role, &mockInterviewPost.Logo, &mockInterviewPost.Description, &mockInterviewPost.Agreement, &mockInterviewPost.PublishedAt)

		mockInterviewPostSlice = append(mockInterviewPostSlice, mockInterviewPost)
	}

	return mockInterviewPostSlice
}
