package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type config struct {
	DBDriver          string
	DbSource          string
	MaxOpenConns      int
	MaxIdleConns      int
	ConnMaxLifetime   time.Duration
	ConnectionTimeout time.Duration
}

type SqlClient struct {
	DB *sql.DB
}

func NewSqlClient(cfg config) (*SqlClient, error) {
	// NewSQLClient creates a new database client with the given configuration.
	db, err := sql.Open(cfg.DBDriver, cfg.DbSource)

	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// Ping the database to verify the connection
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)

	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return &SqlClient{db}, nil
}

// Close terminates the database connection.
func (client *SqlClient) Close() error {
	if client.DB != nil {
		return client.DB.Close()
	}

	return nil
}
