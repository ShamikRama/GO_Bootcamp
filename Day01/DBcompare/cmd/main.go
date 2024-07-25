package main

import (
	pkg_c "Day01/pkg/dbcompare"
	pkg_r "Day01/pkg/dbreader"
	"flag"
)

func main() {
	var data1file, data2file string
	flag.StringVar(&data1file, "old", "", "First database file name")
	flag.StringVar(&data2file, "new", "", "Second database file name")
	flag.Parse()
	old, _ := pkg_r.FileformatRead(data1file)
	new, _ := pkg_r.FileformatRead(data2file)
	pkg_c.Compare(old, new)

}
