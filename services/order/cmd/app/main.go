package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"order/internal/handler"
	"order/internal/order"
	"order/internal/repository"
	"order/internal/service"

	"github.com/dmitryavdonin/gtools/migrations"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//db migrations
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"))

	migrate, err := migrations.NewMigrations(dsn, "file://migrations")
	if err != nil {
		logrus.Fatalf("migrations error: %s", err.Error())
	}

	err = migrate.Up()
	if err != nil {
		logrus.Fatalf("migrations error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: viper.GetString("db.password"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db,
		viper.GetString("api.storeuri"),
		viper.GetString("api.paymenturi"),
		viper.GetString("api.bookuri"),
	)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	var port = viper.GetString("port")

	srv := new(order.Server)
	go func() {
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("order app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("order app shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
