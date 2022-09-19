# SQL Migration for Go Lang

## Usage Example for the SQL migration tool in Go Lang

This project uses the **Migrate** tool. It's a tool written in Go that able to run SQL migrations for multiple databases.

Migrate link [Migrate](https://github.com/golang-migrate/migrate).

This works by managing a table (schema_migrations) allowing you to check which migrations have been applied.

## Starting the Postgres by Docker/Docker-Compose

For better execution of this project, you can run the postgres container with docker and docker-compose. Use the following command in root directory:

```
docker-compose up -d
```

## Migration CLI Instalations

To install the **Migrate CLI** use:

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Creating Migrations

With the **Migrate CLI** installed, you can use it to create the SQL migration files as the folow:

```
migrate create -ext sql -dir db/migration -seq init_schema
```

This code creates the files:

* db/migrations/000001_init_schema.down.sql
* db/migrations/000001_init_schema.up.sql

The files with suffix **up** is used to apply changes and with suffix **down** is used for change reversions

## Running SQL Migrations

Use the parameter **up** or **down** to apply ou revert SQL changes.

```
migrate -path ./src/db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=db_schema" -verbose up
```

## Running SQL Migrations into Go projects

In this example was used the Go 1.17 version and the packages:

* github.com/golang-migrate/migrate
* github.com/golang-migrate/migrate/v4
* github.com/jmoiron/sqlx
* github.com/joho/godotenv

To run migrations in Go lang look to **migrations.go** file that execute the up migration.

[migrations.go](./src/db/migrations.go)

The call database.RunMigrations() in main.go allow to apply migrations automatically in boot application.
