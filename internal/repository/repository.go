package repository

import (
	"context"
	"database/sql"

	"github.com/noolingo/card-service/internal/domain"
)

type Repository interface {
	GetCardByID(ctx context.Context, id string) (*domain.Card, error)
	GetCardByEng(ctx context.Context, eng string) (*domain.Card, error)
}

func New(db *sql.DB) Repository {
	return &card{db: db}
}
