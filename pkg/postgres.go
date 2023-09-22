package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Init postgres config struct
type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// Create new connection
func NewPostgresDB(PostgresConfig PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		PostgresConfig.Host,
		PostgresConfig.Port,
		PostgresConfig.Username,
		PostgresConfig.DBName,
		PostgresConfig.Password,
		PostgresConfig.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
