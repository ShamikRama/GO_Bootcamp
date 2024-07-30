package pkg

import (
	conf "FindPath02/ex01/config"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func ReadFile(filename string, fl conf.Flags) {
	if fl.WordsCount {
		Wordcount(filename)
	}
	if fl.LinesCount {
		Linecount(filename)
	}
	if fl.CharactersCount {
		Characterscount(filename)
	}
}

func Wordcount(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	wordCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println(wordCount)
	return nil
}

func Linecount(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	fmt.Println(lineCount)
	return nil
}

func Characterscount(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	charCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, val := range line {
			if isLetter(val) {
				charCount++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	fmt.Println(charCount)
	return nil

}

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}
