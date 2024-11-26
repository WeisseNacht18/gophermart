package config

import (
	"flag"
	"net/url"
	"os"

	"github.com/WeisseNacht18/gophermart/internal/validator"
)

const (
	defaultRunAddress = "localhost:8080"
)

type Config struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseUri          string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func setValue(dst *string, src string) {
	if src != "" {
		*dst = src
	}
}

func NewConfig() Config {
	result := Config{}

	flag.StringVar(&result.RunAddress, "a", defaultRunAddress, "input run address")
	flag.StringVar(&result.DatabaseUri, "d", "", "input database uri for connecting to database")
	flag.StringVar(&result.AccrualSystemAddress, "r", "", "input accrual system address")

	flag.Parse()

	if validator.IsValidRunAddress(result.RunAddress) != nil {
		result.RunAddress = defaultRunAddress
	}

	if _, err := url.Parse(result.DatabaseUri); err != nil {
		setValue(&result.DatabaseUri, "")
	}

	if _, err := url.Parse(result.AccrualSystemAddress); err != nil {
		setValue(&result.AccrualSystemAddress, "")
	}

	envRunAddress := os.Getenv("RUN_ADDRESS")
	if envRunAddress != "" && validator.IsValidRunAddress(envRunAddress) == nil {
		result.RunAddress = envRunAddress
	}

	envDatabaseUri := os.Getenv("DATABASE_URI")
	if _, err := url.Parse(result.DatabaseUri); envDatabaseUri != "" && err != nil {
		result.DatabaseUri = envDatabaseUri
	}

	envAccrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	if _, err := url.Parse(envAccrualSystemAddress); envAccrualSystemAddress != "" && err == nil {
		result.AccrualSystemAddress = envAccrualSystemAddress
	}

	return result
}
