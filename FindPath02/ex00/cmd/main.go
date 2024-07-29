package main

import (
	"FindPath02/ex00"
	"FindPath02/pkg"
	"log"
)

func main() {
	fl, arg, err := pkg.FindFlagAndArg()
	if err != nil {
		log.Fatal(err)
	}

	if err := ex00.IterateOverEntities(arg, fl); err != nil {
		log.Fatal(err)
	}
}
