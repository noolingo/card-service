package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/noolingo/card-service/internal/domain"
	trans "github.com/noolingo/yandex-dictionary"
)

type card struct {
	db *sql.DB
}

func (c *card) GetCardByID(ctx context.Context, id ...string) ([]*domain.Card, error) {
	args := make([]interface{}, len(id))
	for i, id := range id {
		args[i] = id
	}
	stmt := `SELECT id, eng, rus, transcription from cards where id in (?` + strings.Repeat(",?", len(args)-1) + `)`
	rows, err := c.db.Query(stmt, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	defer rows.Close()

	var cards []*domain.Card
	for rows.Next() {
		card := &domain.Card{}
		err = rows.Scan(&card.ID, &card.Eng, &card.Rus, &card.Transcription)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, err
}

func (c *card) GetCardByEng(ctx context.Context, eng string) (*domain.Card, error) {
	card := &domain.Card{}
	err := c.db.QueryRowContext(ctx, "select * from cards where eng=?", eng).Scan(
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
	err := c.db.QueryRowContext(ctx, "select * from cards where rus=?", rus).Scan(
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

func (c *card) SaveCard(ctx context.Context, translate *trans.Translate) (string, error) {
	ins, err := c.db.PrepareContext(ctx,
		"insert into cards(eng,rus,Transcription)values (?,?,?)")
	if err != nil {
		return "", err
	}
	res, err := ins.ExecContext(ctx, translate.Eng, translate.Rus, translate.Transcription)
	if err != nil {
		return "", err
	}
	id, _ := res.LastInsertId()
	return fmt.Sprint(id), nil
}
