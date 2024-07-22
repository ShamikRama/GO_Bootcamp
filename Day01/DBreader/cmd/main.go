package main

import (
	"Day01/pkg"
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
