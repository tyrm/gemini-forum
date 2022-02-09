package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/juju/loggo"
	// pq used by sqlx
	_ "github.com/lib/pq"
	"github.com/markbates/pkger"
	"github.com/rubenv/sql-migrate"
	"github.com/tyrm/gemini-forum/config"
)

var logger = loggo.GetLogger("db.pq")

// Client is a database client.
type Client struct {
	db *sqlx.DB
}

// NewClient creates a new models Client from Config
func NewClient(cfg *config.Config) (*Client, error) {
	client, err := sqlx.Connect("postgres", cfg.PostgresDsn)
	if err != nil {
		return nil, err
	}

	logger.Debugf("Loading Migrations")
	migrations := &migrate.HttpFileSystemMigrationSource{
		FileSystem: pkger.Dir("/db/postgres/migrations"),
	}

	logger.Debugf("Applying Migrations")
	n, err := migrate.Exec(client.DB, "postgres", migrations, migrate.Up)
	if n > 0 {
		logger.Infof("Applied %d migrations!", n)
	}
	if err != nil {
		logger.Criticalf("Could not migrate models: %s", err)
		return nil, err
	}

	return &Client{
		db: client,
	}, nil
}
