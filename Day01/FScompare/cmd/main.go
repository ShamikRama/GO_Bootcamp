package main

import (
	pkg_s "Day01/pkg/fscompare"
	"flag"
)

func main() {
	var snap1, snap2 string
	flag.StringVar(&snap1, "old", "", "First snap file name")
	flag.StringVar(&snap2, "new", "", "Second snap file name")
	flag.Parse()
	old, _ := pkg_s.GetFiles(snap1)
	new, _ := pkg_s.GetFiles(snap2)
	pkg_s.Compare(old, new)
}
