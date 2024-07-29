package config

import (
	"errors"
	"flag"
)

type Flags struct {
	LinesCount      bool
	CharactersCount bool
	WordsCount      bool
}

var NoFileInput = errors.New("Not File Input")
var FreeFlagsInput = errors.New("You can input >1 flags one time")

func PullArgs(fl Flags, filenames *[]string) error {
	flag.BoolVar(&fl.LinesCount, "l", false, "Print Count of Lines")
	flag.BoolVar(&fl.CharactersCount, "m", false, "Print Count of Characters")
	flag.BoolVar(&fl.WordsCount, "w", false, "Print Count of words")
	flag.Parse()

	*filenames = flag.Args()

	if len(*filenames) == 0 {
		return NoFileInput
	}

	if (fl.CharactersCount || fl.WordsCount) && fl.LinesCount {
		return FreeFlagsInput
	}

	if fl.CharactersCount && fl.WordsCount {
		return FreeFlagsInput
	}

	if !fl.CharactersCount && !fl.LinesCount {
		fl.WordsCount = true
	}

	return nil

}
