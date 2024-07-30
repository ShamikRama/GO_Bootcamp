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

var ErrNoFileInput = errors.New("no file input")
var ErrFreeFlagsInput = errors.New("you can input only one flag at a time")

func PullArgs(fl *Flags, filenames *[]string) error {
	flag.BoolVar(&fl.LinesCount, "l", false, "Print count of lines")
	flag.BoolVar(&fl.CharactersCount, "m", false, "Print count of characters")
	flag.BoolVar(&fl.WordsCount, "w", false, "Print count of words")
	flag.Parse()

	*filenames = flag.Args()

	if len(*filenames) == 0 {
		return ErrNoFileInput
	}

	if (fl.CharactersCount && fl.LinesCount) || (fl.CharactersCount && fl.WordsCount) || (fl.LinesCount && fl.WordsCount) {
		return ErrFreeFlagsInput
	}

	if !fl.CharactersCount && !fl.LinesCount && !fl.WordsCount {
		fl.WordsCount = true
	}

	return nil
}
