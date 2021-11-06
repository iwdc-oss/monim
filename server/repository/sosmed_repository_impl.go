package repository

import (
	"context"
	"database/sql"
	"github.com/iwdc-oss/monim/server/helper"
	"github.com/iwdc-oss/monim/server/model/domain"
)

type SosmedRepositoryImpl struct{}

func NewSosmedRepositoryImpl() *SosmedRepositoryImpl {
	return &SosmedRepositoryImpl{}
}

func (repository *SosmedRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, sosmed domain.Sosmed) domain.Sosmed {
	SQL := "INSERT INTO sosmed(instagram, facebook, twitter, linkedin, github) VALUES(?, ?, ?, ?, ?)"
	res, err := tx.ExecContext(ctx, SQL, sosmed.Instagram, sosmed.Facebook, sosmed.Twitter, sosmed.Linkedin, sosmed.Github)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	sosmed.Id = int32(id)

	return sosmed
}

func (repository *SosmedRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, sosmed domain.Sosmed) domain.Sosmed {
	SQL := "UPDATE sosmed SET instagram = ?, facebook = ?, twitter = ?, linkedin = ?, github = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, sosmed.Instagram, sosmed.Facebook, sosmed.Twitter, sosmed.Github)
	helper.PanicIfError(err)

	return sosmed
}

func (repository *SosmedRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int32) domain.Sosmed {
	SQL := "SELECT instagram, facebook, twitter, linkedin, github FROM sosmed WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	sosmed := domain.Sosmed{}
	rows.Scan(&sosmed.Instagram, &sosmed.Facebook, &sosmed.Twitter, &sosmed.Linkedin, &sosmed.Github)
	sosmed.Id = id

	return sosmed
}
