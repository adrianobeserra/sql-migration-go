package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
	postgresSettings := PostgresDatabaseSettings{}
	dbx, err := postgresSettings.GetPostgresConnection()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return err
	}

	driver, err := postgres.WithInstance(dbx.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return err
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return err
	}

	err = migrator.Up()
	if err != nil {
		fmt.Println("error on execute migrations")
		return err
	}
	log.Print("Migrations executed with success")

	return err
}
