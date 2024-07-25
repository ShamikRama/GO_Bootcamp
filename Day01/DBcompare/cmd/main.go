package main

import (
	pkg "Day01/pkg/dbreader"
	"flag"
)

func main() {
	var data1file, data2file string
	flag.StringVar(&data1file, "data1", "", "First database file name")
	flag.StringVar(&data2file, "data2", "", "Second database file name")
	flag.Parse()
	data1, _ := pkg.FileformatRead(data1file)
	data2, _ := pkg.FileformatRead(data2file)
}
