package main

import (
	pkg "Day01/pkg/dbreader"
	"flag"
	"log"
)

func main() {
	filename := ""
	flag.StringVar(&filename, "f", "", "DataBase file name")
	flag.Parse()
	if filename == "" {
		log.Fatal()
	}
	pkg.ConvertFile(filename)
}
