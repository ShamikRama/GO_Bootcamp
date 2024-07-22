package pkg

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type DBreader interface {
	Read(filename string) (Recip, error)
}

type Conver interface {
	Convert(rec Recip)
}

type Ingredients struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count int    `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

type Cake struct {
	Name       string        `json:"name" xml:"name"`
	Time       int           `json:"time" xml:"stovetime"`
	Ingredient []Ingredients `json:"ingredients" xml:"ingredients>item"`
}

type Recip struct {
	XMLName xml.Name `xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type Json struct {
}

type Xml struct {
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
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return rec, err
	}
	err = json.Unmarshal(content, &rec)
	if err != nil {
		log.Fatal(err)
	}
	return rec, err
}

func (x *Xml) Read(filename string) (rec Recip, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return rec, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(content, &rec)
	if err != nil {
		log.Fatal(err)
	}
	return rec, err
}

func (conv *Xmlconvert) Convert(rec Recip) {
	data, err := json.MarshalIndent(rec, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func (conv *Jsconvert) Convert(rec Recip) {
	data, err := xml.MarshalIndent(rec, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
