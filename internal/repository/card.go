package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/noolingo/card-service/internal/domain"
)

type card struct {
	db *sql.DB
}

func (c *card) GetCardByID(ctx context.Context, id string) (*domain.Card, error) {
	card := &domain.Card{}
	err := c.db.QueryRowContext(ctx, "select * from card where id=?", id).Scan(
		&card.ID,
		&card.ENG,
		&card.RUS,
		&card.EXAMPLE,
		&card.TRANSCRIPTION,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return card, err
}

func (c *card) GetCardByEng(ctx context.Context, eng string) (*domain.Card, error) {
	card := &domain.Card{}
	err := c.db.QueryRowContext(ctx, "select * from card where eng=?", eng).Scan(
		&card.ID,
		&card.ENG,
		&card.RUS,
		&card.EXAMPLE,
		&card.TRANSCRIPTION,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return card, err
}

func (c *card) GetCardByRus(ctx context.Context, rus string) (*domain.Card, error) {
	card := &domain.Card{}
	err := c.db.QueryRowContext(ctx, "select * from card where rus=?", rus).Scan(
		&card.ID,
		&card.ENG,
		&card.RUS,
		&card.EXAMPLE,
		&card.TRANSCRIPTION,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return card, err
}
