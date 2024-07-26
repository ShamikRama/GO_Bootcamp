package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

}
