package sqlite

import (
	"app/internal/storage/internal/migrator"
	"app/internal/storage/sqlite/internal/migration"
	"app/system"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sqlx.DB
}

// Connect - возвращает соединение с хранилищем данных
func Connect(ctx context.Context, dataSourceName string) (*Database, error) {
	db, err := sqlx.Open("sqlite3", dataSourceName+"?_fk=1")
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	return &Database{db: db}, nil
}

// MigrateAll - производит миграции данных
func (storage *Database) MigrateAll(ctx context.Context) error {
	// TODO: использовать другой логгер (сделать его в системном пакете)
	return migrator.MigrateAll(ctx, migration.Migrations, storage.db, true, migrator.Sqlite3Provider)
}

func isApplyWithErr(r sql.Result) (bool, error) {
	c, err := r.RowsAffected()
	if err != nil {
		return false, nil
	}

	return c != 0, nil
}

func isApply(ctx context.Context, r sql.Result) bool {
	apply, err := isApplyWithErr(r)

	system.IfErr(ctx, err)

	return apply
}