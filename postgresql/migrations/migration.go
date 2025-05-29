package migrations

import (
	"database/sql"
	"log"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var once sync.Once

func MigrateUp(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./postgresql/migrations",
		"postgres", driver,
	)
	if err != nil {
		panic(err)
	}
	// golang-migrateが未適用のmigrationのみ実行する
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}
}

func MigrateDown(db *sql.DB) {
	once.Do(func() {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			panic(err)
		}
		m, err := migrate.NewWithDatabaseInstance(
			"file://./postgresql/migrations",
			"postgres", driver,
		)
		if err != nil {
			panic(err)
		}

		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}
	})
}
