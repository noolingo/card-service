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

func (c CardServer) SearchByID(ctx context.Context, req *noolingo.SearchByIDRequest) (*noolingo.SearchReply, error) {
	cards, err := c.service.Card.GetCardByID(ctx, req.Id...)
	resp, _ := newResponse(err)
	var result []*noolingo.CardObject
	for _, card := range cards {
		result = append(result, &noolingo.CardObject{
			Id:            card.ID,
			Eng:           card.Eng,
			Rus:           card.Rus,
			Transcription: card.Transcription,
		})
	}

	return &noolingo.SearchReply{Cards: result, Response: resp}, err
}

func (c CardServer) SearchByRus(ctx context.Context, req *noolingo.SearchByRusRequest) (*noolingo.SearchReply, error) {
	card, err := c.service.Card.GetCardByRus(ctx, req.Rus)
	resp, _ := newResponse(err)
	if err != nil {
		return &noolingo.SearchReply{Cards: nil, Response: resp}, err
	}
	var result []*noolingo.CardObject
	result = append(result, &noolingo.CardObject{
		Id:            card.ID,
		Eng:           card.Eng,
		Rus:           card.Rus,
		Transcription: card.Transcription,
	})
	return &noolingo.SearchReply{Cards: result, Response: resp}, err
}

func (c CardServer) SearchByEng(ctx context.Context, req *noolingo.SearchByEngRequest) (*noolingo.SearchReply, error) {
	card, err := c.service.Card.GetCardByEng(ctx, req.Eng)
	resp, _ := newResponse(err)
	if err != nil {
		return &noolingo.SearchReply{Cards: nil, Response: resp}, err
	}
	var result []*noolingo.CardObject
	result = append(result, &noolingo.CardObject{
		Id:            card.ID,
		Eng:           card.Eng,
		Rus:           card.Rus,
		Transcription: card.Transcription,
	})
	return &noolingo.SearchReply{Cards: result, Response: resp}, err
}
