package pkg

import (
	conf "FindPath02/ex01/config"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string, fl conf.Flags) {
	if fl.WordsCount {
		if err := Wordcount(filename); err != nil {
			fmt.Println(err)
		}
	}
	if fl.LinesCount {
		if err := Linecount(filename); err != nil {
			fmt.Println(err)
		}
	}
	if fl.CharactersCount {
		if err := Characterscount(filename); err != nil {
			fmt.Println(err)
		}
	}
}

func Wordcount(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
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
		return err
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
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	charCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		charCount += len(line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	fmt.Println(charCount)
	return nil
}

/*func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}*/
