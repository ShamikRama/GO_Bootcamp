package internal

import (
	"encoding/csv"
	"log"
	"os"
)

type ReadDB interface {
	ReadDB()
}

// you can add another types and methods
type CSV struct {
}

// function for reading DataBase
func (c *CSV) ReadDB(filepath string) ([][]string, error) {
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
		if err.Error() == "EOF" {
			break
		}
		buf = append(buf, rec)
	}
	return buf, nil
}
