package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DBReader interface {
	Read(filePath string) (Schema, error)
}

type JsonReader struct{}

type PropertiesType struct {
	Type string `json:"type"`
}

type Properties struct {
	Address  PropertiesType `json:"address"`
	Id       PropertiesType `json:"id"`
	Location PropertiesType `json:"location"`
	Name     PropertiesType `json:"name"`
	Phone    PropertiesType `json:"phone"`
}

type Schema struct {
	Properties Properties `json:"properties"`
}

type Location struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type Place struct {
	Address  string   `json:"address"`
	Id       int      `json:"id"`
	Location Location `json:"location"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
}

func (j *JsonReader) Read(filePath string) (Schema, error) {
	b := getDataFromFile(filePath)
	var r Schema

	if !json.Valid(b) {
		err := fmt.Errorf("Not valid JSON")
		return Schema{}, err
	}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return Schema{}, err
	}

	return r, nil
}

func readDB(filePath string, reader DBReader) Schema {
	res, err := reader.Read(filePath)
	if err != nil {
		log.Fatalln("Error reading file.", err)
	}
	return res
}

func getDataFromFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("File does not exists")
		os.Exit(1)
	}
	return data
}

func mappingIndex(indexName string) error {
	// Marshal json file
	schema := readDB("schema.json", new(JsonReader))

	// // Вывести считанный объект в консоль для проверки
	fmt.Printf("Считанный объект Schema: %+v\n", schema)
	return nil
}

func main() {
	schema := readDB("schema.json", new(JsonReader))
	fmt.Printf("Считанный объект Schema: %+v\n", schema)
}
