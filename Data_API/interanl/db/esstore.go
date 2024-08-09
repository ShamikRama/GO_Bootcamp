package db

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type ESstore struct {
	client *elasticsearch.Client
}

func CreateIndex() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	index := "places"

}
