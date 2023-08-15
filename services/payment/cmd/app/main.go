package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"payment/internal/config"
	"payment/internal/handler"
	"payment/internal/payment"
	"payment/internal/repository"
	"payment/internal/service"

	"github.com/dmitryavdonin/gtools/migrations"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.InitConfig("")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	//db migrations
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBname)

	migrate, err := migrations.NewMigrations(dsn, "file://migrations")
	if err != nil {
		logrus.Fatalf("migrations error: %s", err.Error())
	}

	err = migrate.Up()
	if err != nil {
		logrus.Fatalf("migrations error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(dsn)

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db, cfg.API.UserURI)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	var port = cfg.App.Port

	srv := new(payment.Server)
	go func() {
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Printf("Service %s app started", cfg.App.ServiceName)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("Service %s app shutting down", cfg.App.ServiceName)

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
