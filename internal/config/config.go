package config

import "os"

type Config struct {
	DatabaseUrl string
}

func Load() *Config {

	return &Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
}
