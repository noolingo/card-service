package app

import (
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/noolingo/card-service/internal/domain"
	"github.com/sirupsen/logrus"
)

var (
	ErrSignalReceived = errors.New("signal received")
)

func Run(config string) error {
	cfg := &domain.Config{}
	if err := cleanenv.ReadConfig(config, cfg); err != nil {
		return err
	}
	parseFlags(cfg)
	log := logrus.New()
	log.Infof("Hello cards!%#v", cfg)
}
