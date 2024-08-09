package db

import (
	m "Data_Api/models"
	"bytes"
	"encoding/json"
	"fmt"
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

func (s *ESstore) IndexPlace(place m.Place) (err error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(place)
	if err != nil {
		return fmt.Errorf("error encoding place: %s", err)
	}
	res, err := s.client.Index("places", &buf)
	if err != nil {
		return fmt.Errorf("error indexing document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.Status())
	}
	return nil
}
