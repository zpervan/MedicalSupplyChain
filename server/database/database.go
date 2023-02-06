package database

import (
	"database/sql"
	"server/core"

	_ "github.com/lib/pq"
)

// Represents the main database structure
type Database struct {
	logger  *core.Logger // Logging functionality
	handler *sql.DB      // Postgres database handler
}

const (
	db_driver = "postgres"
)

func (d *Database) Connect(psqlInfo string) (*sql.DB, error) {
	d.logger.Info("logging to database with connection string: " + psqlInfo)

	database, err := sql.Open(db_driver, psqlInfo)

	if err != nil {
		return nil, err
	}

	d.logger.Info("database connection established")

	return database, nil
}

func NewDatabase(logger *core.Logger, psqlInfo string) (*Database, error) {
	database := &Database{}
	database.logger = logger

	handler, err := database.Connect(psqlInfo)
	if err != nil {
		return nil, err
	}

	database.handler = handler

	return database, nil
}
