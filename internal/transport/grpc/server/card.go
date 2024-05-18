package grpcserver

import (
	"context"

	"github.com/noolingo/card-service/internal/service"
	"github.com/noolingo/proto/codegen/go/common"
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

func newResponse(err error) (*common.Response, error) {
	response := &common.Response{
		Result: err == nil,
	}

	if err != nil {
		response.Error = &common.Error{
			Error: err.Error(),
		}
	}

	return response, err
}

func (c CardServer) GetCardByID(ctx context.Context, req *noolingo.SearchByIDRequest) (*common.Response, error) {
	card, err := c.service.Card.GetCardByID(ctx, req.Id)
}
