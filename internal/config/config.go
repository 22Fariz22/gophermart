package config

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	runAddress           string
	databaseURI          string
	accrualSystemAddress string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.runAddress, "a", "localhost:8080", "server address")
	flag.StringVar(&cfg.databaseURI, "d", "postgres://postgres:55555@127.0.0.1:5432/gophermart", "database address")
	flag.StringVar(&cfg.accrualSystemAddress, "r", "http://127.0.0.1:8080", "accrual system")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	return cfg, nil
}
