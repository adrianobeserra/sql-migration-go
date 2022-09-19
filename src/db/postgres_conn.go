package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/jmoiron/sqlx"
	//_ "github.com/lib/pq"
)

type PostgresDatabaseSettings struct {
	Host   string `env:"POSTGRES_HOST,required"`
	Port   string `env:"POSTGRES_PORT,required"`
	DbName string `env:"POSTGRES_DB_NAME,required"`
	Schema string `env:"POSTGRES_SCHEMA,required"`
	User   string `env:"POSTGRES_USER,required"`
	Pwd    string `env:"POSTGRES_PWD,required"`
}

func (ref PostgresDatabaseSettings) GetPostgresConnection() (*sqlx.DB, error) {

	fmt.Print(ref.getStringConnection())
	db, err := sqlx.Connect("postgres", ref.getStringConnection())
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return nil, err
	}

	return db, err
}

func (ref PostgresDatabaseSettings) getStringConnection() string {
	dBsettings := loadDotEnvConfig()
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dBsettings.User, strings.Trim(dBsettings.Pwd, "\n"), dBsettings.Host, dBsettings.Port, dBsettings.DbName)
}

func loadDotEnvConfig() PostgresDatabaseSettings {
	err := godotenv.Load("../local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	postgresDatabaseSettings := PostgresDatabaseSettings{}
	postgresDatabaseSettings.Host = os.Getenv("POSTGRES_HOST")
	postgresDatabaseSettings.Port = os.Getenv("POSTGRES_PORT")
	postgresDatabaseSettings.DbName = os.Getenv("POSTGRES_DB_NAME")
	postgresDatabaseSettings.Schema = os.Getenv("POSTGRES_SCHEMA")
	postgresDatabaseSettings.User = os.Getenv("POSTGRES_USER")
	postgresDatabaseSettings.Pwd = os.Getenv("POSTGRES_PWD")
	return postgresDatabaseSettings
}
