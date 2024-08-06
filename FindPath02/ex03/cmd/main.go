package main

import (
	"FindPath02/ex03/config"
	"FindPath02/ex03/pkg"
	"log"
	"sync"
)

func main() {
	var directory string
	var filenames []string
	err := config.FindArgs(&directory, &filenames)
	if err != nil {
		log.Panic(err)
	}
	var wg sync.WaitGroup
	for _, val := range filenames {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			pkg.CompressFile(directory, val)
		}(val)
	}
	wg.Wait()
}
