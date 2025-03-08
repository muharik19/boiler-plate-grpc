package database

import (
	"context"
	"database/sql"
	"time"

	role "github.com/muharik19/boiler-plate-grpc/internal/domain/entities/role"
	"github.com/muharik19/boiler-plate-grpc/pkg/logger"
	global "github.com/muharik19/boiler-plate-grpc/pkg/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var DbBun *bun.DB

func InitBun(ctx context.Context) (*bun.DB, error) {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithDSN(*global.Getenv("BUN_CONNECTION")),
		pgdriver.WithApplicationName(*global.Getenv("APP_NAME")),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	sqldb := sql.OpenDB(pgconn)

	// Set connection pooling parameters
	sqldb.SetMaxOpenConns(25)                  // Maximum open connections
	sqldb.SetMaxIdleConns(10)                  // Maximum idle connections
	sqldb.SetConnMaxLifetime(30 * time.Minute) // Maximum connection life time
	sqldb.SetConnMaxIdleTime(5 * time.Minute)  // Idle time before connection is closed

	db := bun.NewDB(sqldb, pgdialect.New())

	// Add a query hook for logging
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true), // true log query sql, false no log query sql
		bundebug.FromEnv(),
	))

	// Ping the database to test the connection
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	// Auto-create tables
	err = autoCreateTable(ctx, db)
	if err != nil {
		return nil, err
	}

	// Connection successful
	logger.Info("Connected to the database with connection pooling")

	DbBun = db
	return DbBun, nil
}

func autoCreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*role.Role)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
