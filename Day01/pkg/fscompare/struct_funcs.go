package pkg

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// get text from file
func GetFiles(filesnap string) (text []string, err error) {
	file, _ := os.Open(filesnap)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filename := scanner.Text()
		text = append(text, filename)
	}
	err = scanner.Err()
	return text, err
}

// compare 1.txt and 2.txt
func Compare(firstsnap, secondsnap []string) {
	for _, val := range firstsnap {
		if !slices.Contains(secondsnap, val) {
			fmt.Print("Added:", val)
		}
	}
	for _, val := range secondsnap {
		if !slices.Contains(firstsnap, val) {
			fmt.Print("Removed:", val)
		}
	}
}
