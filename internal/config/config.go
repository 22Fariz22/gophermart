package config

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	runAddress           string `env:"RUN_ADDRESS"` //envDefault:":8080"
	databaseURI          string `env:"DATABASE_URI"`
	accrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.runAddress, "a", "localhost:8080", "server address")
	flag.StringVar(&cfg.databaseURI, "d", "", "database address")        //postgres://postgres:55555@127.0.0.1:5432/gophermart
	flag.StringVar(&cfg.accrualSystemAddress, "r", "", "accrual system") // http://127.0.0.1:8080

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	return cfg, nil
}
