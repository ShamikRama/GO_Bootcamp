package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	i "Data_Api/internal"

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
		fmt.Println("Error marshalling into shemabytes")
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

// загрузка данных в es
func pushIntoES(indexname string, idx int, line []string, es *elasticsearch.Client) (err error) {
	id := strconv.Itoa(idx)

	longitude, err := strconv.ParseFloat(line[4], 64)
	if err != nil {
		fmt.Printf("Error parsing to float")
	}

	latitude, err := strconv.ParseFloat(line[5], 64)
	if err != nil {
		fmt.Printf("Error parsing to float")
	}

	place := i.Place{
		Address: line[2],
		Id:      idx,
		Location: i.Location{
			Longitude: longitude,
			Latitude:  latitude,
		},
		Name:  line[1],
		Phone: line[3],
	}

	data, err := json.Marshal(place)
	if err != nil {
		return err
	}

	resp, err := es.Index(indexname, strings.NewReader(string(data)), es.Index.WithDocumentID(id))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}

// заполнение базы
func fillESDATA(indexname string, es *elasticsearch.Client, csvpath string) (err error) {
	data := i.CSVReader{}
	lines, err := data.ReadDB(csvpath)
	if err != nil {
		return nil
	}
	var wg sync.WaitGroup

	for indx, val := range lines[1:] {
		wg.Add(1)
		id := indx
		line := val
		go func() {
			defer wg.Done()
			err = pushIntoES(indexname, id, line, es)
			if err != nil {
				fmt.Println("Error push data inro elastic:", err)
			}

		}()
	}
	wg.Wait()
	return nil
}
