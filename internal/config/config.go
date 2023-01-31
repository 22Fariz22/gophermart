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
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.RUN_ADDRESS, "a", "localhost:8080", "server address")
	flag.StringVar(&cfg.DATABASE_URI, "d", "postgres://postg...", "database address")
	flag.StringVar(&cfg.ACCRUAL_SYSTEM_ADDRESS, "r", "cmd/accrual/accrual_darwin_arm64", "balance system")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	return cfg, nil
}
