package client

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type ESclient struct {
	client *elasticsearch.Client
}

// эту функцию надо подбить под наши условия
func (c *ESclient) Newclient() *elasticsearch.Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error conecting %s", err)
	}

	info, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting information %s", err)
	}
	defer info.Body.Close()
	log.Println(info)
	return es
}
