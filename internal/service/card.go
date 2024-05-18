package service

import (
	"context"

	"github.com/noolingo/card-service/internal/domain"
	"github.com/noolingo/card-service/internal/repository"
	trans "github.com/noolingo/yandex-dictionary"
	"github.com/sirupsen/logrus"
)

type CardService struct {
	logger     *logrus.Logger
	Config     *domain.Config
	repository *repository.Repository
}

func NewCardService(p *Param) *CardService {
	return &CardService{
		logger:     p.Logger,
		repository: p.Repository,
		Config:     p.Config,
	}
}

//getCard() x3(id, rus, eng)

func (c *CardService) SaveCardRus(ctx context.Context, wordRus string, api string) (err error) {
	translate, err := trans.TranslateRus(wordRus, api)
	if err != nil {
		return err
	}

	return err
}
