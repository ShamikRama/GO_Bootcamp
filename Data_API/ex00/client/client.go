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

	"github.com/elastic/go-elasticsearch/v7"
)

const (
	shemapath string = "shema.json"
	url       string = "http://localhost:9200/"
)

// функция создания клиента
func GetConnectionClient() *elasticsearch.Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
		},
		// Timeout: 30 * time.Second, // Увеличьте время ожидания
	}
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{url},
		Transport: httpClient.Transport,
	})
	if err != nil {
		log.Fatalf("Error creating client %s", err)
	}
	return es
}

// создание индекса
func CreateIndex(indexName string, es *elasticsearch.Client) error {
	_, err := es.Indices.Delete([]string{indexName})
	if err != nil {
		return fmt.Errorf("error deleting index: %w", err)
	}
	res, err := es.Indices.Create(indexName)
	if err != nil {
		return fmt.Errorf("error creating index: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("failed to create index: %s", res.String())
	}
	return nil
}

// маппинг то есть создание структуры как будут храниться данные
func Mapping(nameIndex string, es *elasticsearch.Client) error {
	shema := i.JsonReader{}
	sh, err := shema.ReadDB(shemapath)
	if err != nil {
		return fmt.Errorf("error reading schema: %w", err)
	}
	sh.Properties.Id.Type = "long"
	shemabytes, err := json.MarshalIndent(sh, "", "  ") // Используем json.MarshalIndent
	if err != nil {
		return fmt.Errorf("error marshalling schema: %w", err)
	}
	req, err := http.NewRequest(http.MethodPut, url+nameIndex+"/place/_mapping", bytes.NewBuffer(shemabytes))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("include_type_name", "true")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// заполнение базы
func FillData(indexname string, es *elasticsearch.Client, csvpath string) error {
	data := i.CSVReader{}
	lines, err := data.ReadDB(csvpath)
	if err != nil {
		return fmt.Errorf("error reading CSV: %w", err)
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	for indx, val := range lines[1:] {
		wg.Add(1)
		id := indx
		line := val
		go func() {
			defer wg.Done()
			err := pushIntoES(indexname, id, line, es)
			if err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	if len(errors) > 0 {
		return fmt.Errorf("errors occurred during data push: %v", errors)
	}

	return nil
}

// загрузка данных в es
func pushIntoES(indexname string, idx int, line []string, es *elasticsearch.Client) error {
	id := strconv.Itoa(idx)

	longitude, err := strconv.ParseFloat(line[4], 64)
	if err != nil {
		return fmt.Errorf("error parsing longitude: %w", err)
	}

	latitude, err := strconv.ParseFloat(line[5], 64)
	if err != nil {
		return fmt.Errorf("error parsing latitude: %w", err)
	}

	place := i.Place{
		Id:      idx,
		Address: line[2],
		Location: i.Location{
			Longitude: longitude,
			Latitude:  latitude,
		},
		Name:  line[1],
		Phone: line[3],
	}

	data, err := json.MarshalIndent(place, "", "  ") // Используем json.MarshalIndent
	if err != nil {
		return fmt.Errorf("error marshalling place: %w", err)
	}

	resp, err := es.Index(indexname, strings.NewReader(string(data)), es.Index.WithDocumentID(id))
	if err != nil {
		return fmt.Errorf("error indexing document ID=%s: %w", id, err)
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return fmt.Errorf("error indexing document ID=%s: %s", id, resp.String())
	}
	return nil
}
