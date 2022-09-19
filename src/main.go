package main

import (
	database "github.com/adrianobeserra/sql-migration-go/db"
)

func main() {
	database.RunMigrations()
}
