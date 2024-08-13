package internal

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// you can add another types and methods
type CSVReader struct{}

type JsonReader struct{}

// function for reading DataBase
func (c *CSVReader) ReadDB(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %s", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	var buf [][]string
	for {
		rec, err := reader.Read()
		if err != nil && err.Error() == "EOF" {
			break
		}
		buf = append(buf, rec)
	}
	return buf, nil
}

func (j *JsonReader) ReadDB(filepath string) (Schema, error) {
	buf, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading the file %s", err)
	}
	var s Schema

	err = json.Unmarshal(buf, &s)
	if err != nil {
		fmt.Println("Ошибка при десериализации данных:", err)
		return Schema{}, err
	}
	return s, nil
}
