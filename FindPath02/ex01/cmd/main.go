package main

import (
	conf "FindPath02/ex01/config"
	pkg "FindPath02/ex01/pkg"
	"fmt"
)

func main() {
	var filenames []string
	var flags conf.Flags
	err := conf.PullArgs(&flags, &filenames)
	if err != nil {
		fmt.Println(err)
	}
	for _, filename := range filenames {
		pkg.ReadFile(filename, flags)
	}
}
