package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func GetFiles(filesnap string) (files []string, err error) {
	file, err := os.Open(filesnap)
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filename := scanner.Text()
		files = append(files, filename)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error:", err)
		return files, err
	}
	return files, err
}

func Compare(firstsnap, secondsnap string) {
	oldfile, err := GetFiles(firstsnap)
	if err != nil {
		fmt.Println("Error getting file")
	}
	newfile, err := GetFiles(secondsnap)
	if err != nil {
		fmt.Println("Error getting file")
	}

	for _, val := range oldfile {
		if !slices.Contains(newfile, val) {
			fmt.Print("Added:", val)
		}
	}
	for _, val := range newfile {
		if !slices.Contains(oldfile, val) {
			fmt.Print("Removed:", val)
		}
	}
}
