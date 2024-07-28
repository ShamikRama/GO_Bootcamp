package main

import (
	m "FindPath02/ex00"
	pkg "FindPath02/pkg"
	"fmt"
)

func main() {
	flag, arg, err := pkg.FindFlagAndArg()
	if err != nil {
		fmt.Println(err)
		return
	}
	m.IterateOverEntities(arg, flag)
}
