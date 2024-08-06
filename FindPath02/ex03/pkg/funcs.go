package pkg

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
)

func CompressFile(filepath string, filenames string) {
	file, err := os.Open(filenames)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	info, err := os.Stat(filenames)
	if err != nil {
		log.Fatal(err)
	}

	fullPath := CreateArchiveName(filepath, filenames, info)

	archive, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	writerGzip, err := gzip.NewWriterLevel(archive, gzip.BestCompression)
	if err != nil {
		log.Fatal(err)
	}
	defer writerGzip.Close()

	writerTar := tar.NewWriter(writerGzip)
	defer writerTar.Close()

	hdr := &tar.Header{
		Name: path.Base(filenames),
		Mode: 0644,
		Size: info.Size(),
	}
	if err := writerTar.WriteHeader(hdr); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(writerTar, file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %s compressed to %s\n", filenames, fullPath)
}

func CreateArchiveName(filepath string, filenames string, info os.FileInfo) string {
	timestamp := strconv.FormatInt(info.ModTime().Unix(), 10)
	if filepath != "" {
		return path.Join(filepath, path.Base(filenames)+"_"+timestamp+".tar.gz")
	}
	return path.Base(filenames) + "_" + timestamp + ".tar.gz"
}
