package mysql

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/noolingo/card-service/internal/domain"
)

func New(cfg *domain.Mysql) (*sql.DB, error) {
	dsn := strings.TrimPrefix(cfg.DSN, "mysql://")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return db, err
}
