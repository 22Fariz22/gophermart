package config

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	RUN_ADDRESS            string
	DATABASE_URI           string
	ACCRUAL_SYSTEM_ADDRESS string
	//LogLevel               string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.RUN_ADDRESS, "a", "localhost:8080", "server address")
	flag.StringVar(&cfg.DATABASE_URI, "d", "postgres://postgres:55555@127.0.0.1:5432/gophermart", "database address")
	flag.StringVar(&cfg.ACCRUAL_SYSTEM_ADDRESS, "r", "http://127.0.0.1:8088", "accrual system")
	//flag.StringVar(&cfg.LogLevel, "l", "debug", "Log level")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	return cfg, nil
}
