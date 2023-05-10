package repository

import (
	"bryanagamk/go-rest-miwt/domain"
	"context"
	"database/sql"
)

type CutiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, cuti domain.Cuti) domain.Cuti
	Update(ctx context.Context, tx *sql.Tx, cuti domain.Cuti) domain.Cuti
	Delete(ctx context.Context, tx *sql.Tx, cuti domain.Cuti)
	FindById(ctx context.Context, tx *sql.Tx, cutiId string) (domain.Cuti, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Cuti
}
