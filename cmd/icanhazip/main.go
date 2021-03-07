package main

import (
	"context"
	"net/http"
	"os"
	"project-template/internal/config"
	"project-template/internal/resource/icanhazip"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	requestTimeout = 10 * time.Second
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

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	resource := icanhazip.NewResource(&icanhazip.Config{
		Endpoint: cfg.Icanhazip.Endpoint,
	})

	output, err := resource.GetIP(ctx, icanhazip.GetIPInput{
		Client: http.DefaultClient,
	})
	if err != nil {
		return errors.Wrap(err, "icanhazip.GetIP")
	}

	os.Stdout.WriteString("Your IP is " + output.IP + "\n")

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
