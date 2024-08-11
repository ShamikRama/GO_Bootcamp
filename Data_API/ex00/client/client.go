package client

import (
	"encoding/csv"
	"log"
	"os"

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

// функция для чтения csv файла и его парсинг
func GetDataFromFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %s", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	var buf [][]string
	for {
		rec, err := reader.Read()
		if err.Error() == "EOF" {
			break
		}
		buf = append(buf, rec)
	}
	return buf, nil
}
