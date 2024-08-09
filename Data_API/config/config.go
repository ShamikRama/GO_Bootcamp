package config

import (
	"os"
)

type Config struct {
	ElasticSearchURL string
}

func LoadConfig() *Config {
	return &Config{
		ElasticSearchURL: os.Getenv("ELASTICSEARCH_URL"),
	}
}
