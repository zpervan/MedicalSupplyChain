package database

import (
	"database/sql"
	"server/core"
	_ "github.com/lib/pq"
)

type Database struct {
	logger *core.Logger
}

const (
	db_driver = "postgres"
)

func (d *Database) Connect(psqlInfo string) error {
	d.logger.Info("logging to database with connection string: " + psqlInfo)

	database, err := sql.Open(db_driver, psqlInfo)

	if err != nil {
		return err
	}

	err = database.Ping()
	if err != nil {
		return err
	}

	d.logger.Info("database connection established")

	return nil
}

func NewDatabase(logger *core.Logger, psqlInfo string) (*Database, error) {
	database := &Database{}
	database.logger = logger

	err := database.Connect(psqlInfo)
	if err != nil {
		return nil, err
	}

	return database, nil
}
