package grpcserver

import (
	"github.com/noolingo/card-service/internal/service"
	"github.com/noolingo/proto/codegen/go/noolingo"
	"github.com/sirupsen/logrus"
)

type CardServer struct {
	noolingo.UnimplementedCardsServer
	logger  *logrus.Logger
	service *service.Services
}

func newCardServer(logger *logrus.Logger, service *service.Services) CardServer {
	return CardServer{logger: logger, service: service}
}
