package config

import (
	"os"
)

type Postgres struct {
	Username      string
	Password      string
	Host          string
	Port          string
	Database      string
	SslMode       string
	MigrationPath string
}

type GeneralConfig struct {
	PG_Config        Postgres
	JsonPath         string
	PasswordFilePath string
}

func SetConfigs(cfg *GeneralConfig) {
	cfg.JsonPath = os.Getenv("JSON_PATH")
	cfg.PasswordFilePath = os.Getenv("PASSWD_PATH")
	// Postgres Database configs
	cfg.PG_Config.Username = os.Getenv("PG_USERNAME")
	cfg.PG_Config.Password = os.Getenv("PG_PASSWORD")
	cfg.PG_Config.Host = os.Getenv("PG_HOST")
	cfg.PG_Config.Port = os.Getenv("PG_PORT")
	cfg.PG_Config.Database = os.Getenv("PG_DATABASE")
	cfg.PG_Config.SslMode = os.Getenv("PG_SSLMODE")
	cfg.PG_Config.MigrationPath = os.Getenv("PG_MIGRATION_PATH")

}
