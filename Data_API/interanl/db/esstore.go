package db

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type ESstore struct {
	client *elasticsearch.Client
}

func NewESstore() (*ESstore, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &ESstore{
		client: client,
	}, nil
}
