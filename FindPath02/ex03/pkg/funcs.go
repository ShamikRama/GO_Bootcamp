package pkg

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

type Fileinfo struct {
	val_1 tar.Header
	val_2 gzip.Header
}

func CompressFile(filepath string, filenames string) {
	file, err := os.Open(filenames)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, err := os.Stat(filenames)
	if err != nil {
		fmt.Println(err)
		return
	}
	fullPath := CreateAchiveName(filepath, filenames, info)
	archive, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archive.Close()
	writter_gzip, err := gzip.NewWriterLevel(archive, gzip.BestCompression)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archive.Close()

	writter_tar := tar.NewWriter(writter_gzip)

}

func CreateAchiveName(filepath1 string, filenames1 string, info1 os.FileInfo) string {
	timestamp := strconv.FormatInt(info1.ModTime().Unix(), 10)
	if filepath1 != "" {
		return path.Join(filepath1, filenames1+"_"+timestamp+".tar.gz")
	} else {
		return path.Base(filenames1) + "_" + timestamp + ".tar.gz"
	}
}
