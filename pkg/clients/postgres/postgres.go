package postgres

import (
	"fmt"

	"github.com/strpc/resume-success/pkg/logging"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	*sqlx.DB
}

func NewClient(host, user, password, dbName, SSLMode string, port int) (*Client, error) {
	logger := logging.GetLogger()
	uriDb := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", user, password, host, port, dbName, SSLMode)
	db, err := sqlx.Connect("pgx", uriDb)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	logger.Infof("Postgresql is connected to %s:%d", host, port)

	if logger.IsDebug() {
		logger.Debugf("%+v", db.Stats())
	}

	return &Client{
		db,
	}, nil
}
