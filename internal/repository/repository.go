package repository

import (
	"context"
	"database/sql"

	"github.com/noolingo/card-service/internal/domain"
	trans "github.com/noolingo/yandex-dictionary"
)

type Repository interface {
	GetCardByID(ctx context.Context, id ...string) ([]*domain.Card, error)
	GetCardByEng(ctx context.Context, eng string) (*domain.Card, error)
	GetCardByRus(ctx context.Context, rus string) (*domain.Card, error)
	SaveCard(ctx context.Context, translate *trans.Translate) (string, error)
}

func New(db *sql.DB) Repository {
	return &card{db: db}
}
