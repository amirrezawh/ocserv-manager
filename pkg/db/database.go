package database

import (
	"database/sql"
	"fmt"

	"github.com/amirrezawh/ocserv-manager/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type Users struct {
	ID         uint `gorm:primaryKey`
	Username   string
	RX_TX_BYTE uint64
	RX_TX      string
	LIMIT      uint64
	Active     bool
}

func ConnectToPostgres(cfg *config.GeneralConfig) *sql.DB {

	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PG_Config.Username,
		cfg.PG_Config.Password,
		cfg.PG_Config.Host,
		cfg.PG_Config.Port,
		cfg.PG_Config.Database,
		cfg.PG_Config.SslMode)

	conn, err := sql.Open("postgres", connectionStr)

	if err != nil {
		panic(err)
	}

	return conn
}

func Migrate(cfg *config.GeneralConfig) {

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	pg := ConnectToPostgres(cfg)

	if err := goose.Up(pg, cfg.PG_Config.MigrationPath); err != nil {
		panic(err)
	}
}
