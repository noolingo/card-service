package service

import (
	"context"
	"errors"

	"github.com/noolingo/card-service/internal/domain"
	"github.com/noolingo/card-service/internal/repository"
	trans "github.com/noolingo/yandex-dictionary"
	"github.com/sirupsen/logrus"
)

var (
	ErrNoCardFound = errors.New("no such card found")
)

type CardService struct {
	logger     *logrus.Logger
	Config     *domain.Config
	repository repository.Repository
}

func NewCardService(p *Params) *CardService {
	return &CardService{
		logger:     p.Logger,
		repository: *p.Repository,
		Config:     p.Config,
	}
}

func (c *CardService) GetCardByID(ctx context.Context, id string) (*domain.Card, error) {
	card, err := c.repository.GetCardByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if card == nil {
		return nil, ErrNoCardFound
	}
	return card, err
}

func (c *CardService) GetCardByRus(ctx context.Context, rus string) (*domain.Card, error) {
	card, err := c.repository.GetCardByRus(ctx, rus)
	if err != nil {
		return nil, err
	}
	if card == nil {
		err := c.saveCardRus(ctx, rus)
		if err != nil {
			return nil, err
		}
	}
	return card, err
}

func (c *CardService) GetCardByEng(ctx context.Context, eng string) (*domain.Card, error) {
	card, err := c.repository.GetCardByEng(ctx, eng)
	if err != nil {
		return nil, err
	}
	if card == nil {
		err := c.saveCardEng(ctx, eng)
		if err != nil {
			return nil, err
		}
	}
	return card, err
}

func (c *CardService) saveCardRus(ctx context.Context, rus string) error {
	translate, err := trans.TranslateRus(rus, c.Config.YandexDictionary.Api)
	if err != nil {
		return err
	}
	_, err = c.repository.SaveCard(ctx, translate)
	if err != nil {
		return err
	}
	return nil
}

func (c *CardService) saveCardEng(ctx context.Context, eng string) error {
	translate, err := trans.TranslateEng(eng, c.Config.YandexDictionary.Api)
	if err != nil {
		return err
	}
	_, err = c.repository.SaveCard(ctx, translate)
	if err != nil {
		return err
	}
	return nil
}
