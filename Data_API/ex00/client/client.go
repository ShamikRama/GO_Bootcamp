package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	i "Data_Api/ex00/internal"

	"github.com/elastic/go-elasticsearch/v8"
)

const (
	shemapath string = "src/ex00/cmd/shema.json"
	url       string = "http://localhost:9200/"
)

// функция создания клиента
func GetConnectionClient() *elasticsearch.Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
		},
	}
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: httpClient.Transport,
	})
	if err != nil {
		log.Fatalf("Error creating client %s", err)
	}
	return es
}

// создание индекса
func CreateIndex(indexName string, es *elasticsearch.Client) (err error) {
	_, err = es.Indices.Delete([]string{indexName})
	if err != nil {
		return err
	}
	res, err := es.Indices.Create(indexName)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("failed to create index: %s", res.String())
	}
	return nil
}

// маппинг то есть создание структуры как будут хранится данные
func Mapping(nameIndex string, es *elasticsearch.Client) error {
	shema := i.JsonReader{}
	sh, err := shema.ReadDB(shemapath)
	if err != nil {
		return err
	}
	sh.Properties.Id.Type = "long"
	shemabytes, err := json.Marshal(sh)
	if err != nil {
		fmt.Println("FIFIFIF")
	}
	req, err := http.NewRequest(http.MethodPut, url+nameIndex+"/place/_mapping", bytes.NewBuffer(shemabytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("include_type_name", "true")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
