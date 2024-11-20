package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// "github.com/Maksim646/Bot/database"
	// _ "github.com/Maksim646/Bot/domain/user/repository/postgresql"

	// _tgBot "github.com/Maksim646/Bot/bot"
	// _userRepo "github.com/Maksim646/Bot/domain/user/repository/postgresql"
	// _userUsecase "github.com/Maksim646/Bot/domain/user/usecase"
	// "github.com/Maksim646/Bot/handler"
	//"github.com/heetch/sqalx"
	"github.com/Maksim646/TestAssignment/api/server/restapi/database/postgresql"
	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

const (
	httpVersion = "development"
)

var config struct {
	Addr          string `envconfig:"ADDR" default:":8095"`
	PostgresURI   string `envconfig:"POSTGRES_URI" default:"postgres://postgres:postgres@localhost:5433/bot_db?sslmode=disable"`
	MigrationsDir string `envconfig:"MIGRATIONS_DIR" default:"database/migrations"`
}

func main() {
	envconfig.MustProcess("", &config)

	fmt.Println("db:", config.PostgresURI)

	sqlxConn, err := sqlx.Connect("postgres", config.PostgresURI)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}

	migrator := postgresql.NewMigrator(config.PostgresURI, config.MigrationsDir)
	if err := migrator.Apply(); err != nil {
		log.Fatal("cannot apply migrations: ", err, config.MigrationsDir)
	}

	defer sqlxConn.Close()

	log.Println("Successfully connected to the database")

	sqlxConn.SetMaxOpenConns(100)
	sqlxConn.SetMaxIdleConns(100)
	sqlxConn.SetConnMaxLifetime(5 * time.Minute)

	sqalxConn, err := sqalx.New(sqlxConn)
	if err != nil {
		log.Fatal("cannot create sqalx connection: ", err)
	}
	defer sqalxConn.Close()

	log.Println("db was initialized")

	userRepo := _userRepo.New(sqalxConn)
	userUsecase := _userUsecase.New(userRepo)

	handler.New(
		httpVersion,
		userUsecase,
	)

	server := http.Server{
		Addr: config.Addr,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

}
