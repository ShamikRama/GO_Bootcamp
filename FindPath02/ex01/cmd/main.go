package main

import (
	conf "FindPath02/ex01/config"
	pkg "FindPath02/ex01/pkg"
	"fmt"
	"sync"
)

func main() {
	var filenames []string
	var flags conf.Flags
	err := conf.PullArgs(&flags, &filenames)
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup
	for _, filename := range filenames {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			pkg.ReadFile(filename, flags)
		}(filename)
	}
	wg.Wait()
}
