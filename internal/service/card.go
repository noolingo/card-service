package service

import (
	"github.com/noolingo/card-service/internal/domain"
	"github.com/noolingo/card-service/internal/repository"
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
