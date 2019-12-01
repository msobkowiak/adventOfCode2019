package reader

import (
	"os"
	"log"
	"bufio"
	"strconv"
)

type FileReader interface {
	Read(fileName string) []int
}

type fileReader struct {}

func NewFileReader() FileReader {
	return &fileReader{}
}

func (r *fileReader) Read(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	data := make([]int, 0, 0)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		data = append(data, value)
	}

	return data
}