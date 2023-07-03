package database

import (
	"database/sql"
	"fmt"

	"github.com/feryadialoi/go-url-shortener/config"
	_ "github.com/lib/pq"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
	"github.com/sirupsen/logrus"
)

func connectionString(conf config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DBPort,
		conf.DBUser,
		conf.DBPass,
		conf.DBName,
	)
}

func New(conf config.Config) (*sql.DB, error) {
	logrus.Info("New Database")

	db, err := sql.Open("nrpostgres", connectionString(conf))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MustNew(conf config.Config) *sql.DB {
	db, err := New(conf)
	if err != nil {
		logrus.Fatalf("Failed to create new database client, err: %v", err)
	}

	return db
}
