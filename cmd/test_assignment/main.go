package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/Maksim646/TestAssignment/internal/api/server/restapi/handler"
	"github.com/Maksim646/TestAssignment/internal/database/postgresql"
	"github.com/Maksim646/TestAssignment/pkg/logger"

	_songRepo "github.com/Maksim646/TestAssignment/internal/domain/song/repository/postgresql"
	_songUsecase "github.com/Maksim646/TestAssignment/internal/domain/song/usecase"
)

const (
	httpVersion = "development"
)

var config struct {
	Addr          string `envconfig:"ADDR" default:"8000"`
	LogLevel      string `envconfig:"LOG_LEVEL"`
	MigrationsDir string `envconfig:"MIGRATIONS_DIR" default:"../../internal/database/postgresql/migrations"`
	PostgresURI   string `envconfig:"POSTGRES_URI" default:"postgres://postgres:12345@localhost:5448/test_assignment_db?sslmode=disable"`

	HashSalt     string `envconfig:"HASH_SALT" default:"MaximAdamov2002"`
	JWTSigninKey string `envconfig:"JWT_SIGNIN_KEY" default:"MaximAdamov2002"`
}

func main() {
	envconfig.MustProcess("", &config)

	if err := logger.BuildLogger(config.LogLevel); err != nil {
		log.Fatal("cannot build logger: ", err)
	}

	zap.L().Info("PostgresURI: ", zap.String("uri", config.PostgresURI))
	zap.L().Info("MigrationsDir: ", zap.String("dir", config.MigrationsDir))

	time.Sleep(3 * time.Second)

	migrator := postgresql.NewMigrator(config.PostgresURI, config.MigrationsDir)
	if err := migrator.Apply(); err != nil {
		log.Fatal("cannot apply migrations: ", err)
	}

	sqlxConn, err := sqlx.Connect("postgres", config.PostgresURI)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}

	sqlxConn.SetMaxOpenConns(100)
	sqlxConn.SetMaxIdleConns(100)
	sqlxConn.SetConnMaxLifetime(5 * time.Minute)

	defer sqlxConn.Close()

	sqalxConn, err := sqalx.New(sqlxConn)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}
	defer sqalxConn.Close()

	zap.L().Info("Database manage was process successfully")

	songRepo := _songRepo.New(sqalxConn)
	songUsecase := _songUsecase.New(songRepo)

	// adminRepo := _adminRepo.New(sqalxConn)
	// adminUsecase := _adminUsecase.New(adminRepo)

	// projectRepo := _projectRepo.New(sqalxConn)
	// projectUsecase := _projectUsecase.New(projectRepo)

	// chassisRepo := _chassisRepo.New(sqalxConn)
	// chassisUsecase := _chassisUsecase.New(chassisRepo)

	appHandler := handler.New(
		// userUsecase,
		// adminUsecase,
		// projectUsecase,
		// chassisUsecase,
		songUsecase,

		httpVersion,
		config.HashSalt,
		config.JWTSigninKey,
	)

	chain := alice.New(appHandler.WsMiddleware).Then(appHandler)
	if chain == nil {
	}
	server := http.Server{
		Handler: chain,
		Addr:    ":" + config.Addr,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Ожидание сигнала завершения (SIGINT, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Блокируем `main`, ожидая сигнала
	<-quit
	zap.L().Info("Shutdown server ...")

	// Создаем таймаут для выключения сервера
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Завершение сервера
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	zap.L().Info("Server exiting")
}
