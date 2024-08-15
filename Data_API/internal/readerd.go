package internal

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type DBReader interface {
	ReadDB(filepath string) (interface{}, error)
}

// you can add another types and methods
type CSVReader struct{}

type JsonReader struct{}

// function for reading DataBase
func (c *CSVReader) ReadDB(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func (j *JsonReader) ReadDB(filepath string) (Schema, error) {
	buf, err := os.ReadFile(filepath)
	if err != nil {
		return Schema{}, fmt.Errorf("error reading the file: %w", err)
	}
	var s Schema

	err = json.Unmarshal(buf, &s)
	if err != nil {
		return Schema{}, fmt.Errorf("ошибка при десериализации данных: %w", err)
	}
	return s, nil
}
