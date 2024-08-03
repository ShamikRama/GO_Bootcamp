package main

import (
	"errors"
	"flag"
	"log"
	"os"
)

var ErrNoDir = errors.New("No directiry")
var ErrNoFile = errors.New("No file input")

func Config() (directory string, filenames []string, err error) {
	flag.StringVar(&directory, "a", "", "Directory to cp")
	flag.Parse()
	if directory == "" {
		log.Panic()
	}
	info, err := os.Stat(directory)
	if err != nil {
		return directory, nil, err
	}
	if !info.IsDir() {
		err = ErrNoDir
		return directory, nil, err
	}
	filenames = flag.Args()
	if filenames == nil {
		err = ErrNoFile
		return directory, nil, err
	}
	return
}
