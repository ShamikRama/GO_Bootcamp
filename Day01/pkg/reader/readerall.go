package pkg

import (
	"encoding/xml"
	"errors"
	"flag"
)

const Nofileerror = "Передайте название файла через флаг -f"

type DBreader interface {
	Readfile(dbfilename string) (Recip, error)
}

type DBConverter interface {
	Convert(cakes Recip)
}

type Recip struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

func Сheckdb() (dbfilename string, err error) {
	flag.StringVar(&dbfilename, "f", "", "DataBase file name")
	flag.Parse()

	if dbfilename == "" {
		return "", errors.New(Nofileerror)
	}
	return dbfilename, err
}
