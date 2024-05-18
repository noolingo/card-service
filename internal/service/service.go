package service

import (
	"context"

	"github.com/noolingo/card-service/internal/domain"
	"github.com/noolingo/card-service/internal/repository"
	"github.com/sirupsen/logrus"
)

type Card interface {
	GetCardByID(ctx context.Context, id string) (*domain.Card, error)
	GetCardByEng(ctx context.Context, eng string) (*domain.Card, error)
	GetCardByRus(ctx context.Context, rus string) (*domain.Card, error)
}

type Services struct {
	Card Card
}

type Params struct {
	Logger     *logrus.Logger
	Config     *domain.Config
	Repository *repository.Repository
}

func New(p *Params) *Services {
	return &Services{
		Card: NewCardService(p),
	}
}
