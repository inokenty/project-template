package main

import (
	"net/http"
	"project-template/internal/config"
	"project-template/internal/repo"
	"project-template/internal/repo/model"
	"project-template/internal/service"
	"project-template/internal/transport/rest/handler"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err).Msg("run")
	}
}

func run() error {
	cfg, err := getConfig()
	if err != nil {
		return errors.Wrap(err, "getConfig")
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		return errors.New("gorm.Open")
	}

	if err := autoMigrate(db); err != nil {
		return errors.Wrap(err, "autoMigrate")
	}

	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	repos := repo.New(db)
	services := service.New(repos)
	handlers := handler.New(services)

	// Add all routes
	handlers.AddRoutes(router)

	log.Info().Int("port", cfg.Server.Port).Msg("ListenAndServe")

	listenAddr := ":" + strconv.Itoa(cfg.Server.Port)

	if err := http.ListenAndServe(listenAddr, router); err != nil {
		return errors.Wrap(err, "http.ListenAndServe")
	}

	return nil
}

func getConfig() (*config.Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	var cfg config.Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return &cfg, nil
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.Order{},
		&model.User{},
	); err != nil {
		return errors.Wrap(err, "db.AutoMigrate")
	}

	return nil
}
