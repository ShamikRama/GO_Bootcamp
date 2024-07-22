package pkg

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"os"
)

type DBreader interface {
	Read(filename string) (Recip, error)
}

type Ingredients struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count int    `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

// здесь возможно вот так `json:"ingredients" xml:"ingredients>item"`
type Cake struct {
	Name       string        `json:"name" xml:"name"`
	Time       int           `json:"time" xml:"stovetime"`
	Ingredient []Ingredients `json:"ingredients" xml:"ingredients"`
}

type Recip struct {
	Cakes []Cake `json:"cake" xml:"cake"`
}

type Json struct {
	cakes Recip
}

type Xml struct {
	cakes Recip
}

type Jsconvert struct {
}

type Xmlconvert struct {
}

func (j *Json) Read(filename string) (rec Recip, err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &rec)
	if err != nil {
		log.Fatal(err)
	}
	return j.cakes, err
}

func (x *Xml) Read(filename string) (rec Recip, err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(content, &rec)
	if err != nil {
		log.Fatal(err)
	}
	return x.cakes, err
}
