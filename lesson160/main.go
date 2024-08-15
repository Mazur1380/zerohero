package main

import (
	"errors"
	"io"
	"log"
	"os"
)

var ErrNotFound = errors.New("file not found")

func main() {

	_, err := readFile("txt.txt")
	// if err != nil {
	// log.Fatal(err)
	// }
	if errors.Is(err, ErrNotFound) {
		log.Print("Файл не  найден")
	}

}

func readFile(name string) (string, error) {
	f, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if errors.Is(err, os.ErrNotExist) {
		return "", ErrNotFound
	} else if err != nil {
		return "", err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	s := string(b)

	return s, nil
}
