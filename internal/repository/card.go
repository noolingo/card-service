package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/noolingo/card-service/internal/domain"
	trans "github.com/noolingo/yandex-dictionary"
)

type card struct {
	db *sql.DB
}

func (c *card) GetCardByID(ctx context.Context, id string) (*domain.Card, error) {
	card := &domain.Card{}
	err := c.db.QueryRowContext(ctx, "select * from card where id=?", id).Scan(
		&card.ID,
		&card.Eng,
		&card.Rus,
		&card.Transcription,
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
		&card.Eng,
		&card.Rus,
		&card.Transcription,
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
		&card.Eng,
		&card.Rus,
		&card.Transcription,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return card, err
}

func (c *card) SaveCard(ctx context.Context, translate *trans.Translate) (id string, err error) {
	ins, err := c.db.PrepareContext(ctx,
		"insert into card(eng,rus,Transcription)values (?,?,?,?)")
	if err != nil {
		return "", err
	}
	res, err := ins.ExecContext(ctx, translate.Eng, translate.Rus, translate.Transcription)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(res.LastInsertId()), nil
}
