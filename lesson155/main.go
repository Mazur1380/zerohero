package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// err := CopyFile("in.txt", "file.txt", "file2.txt")
	// if err != nil {
	// log.Fatal(err)
	// }

	d := FindWord("in.txt", "руки")
	if d {
		fmt.Println("Слово найдено")

	} else {
		fmt.Println("Слово не найдено")
	}

}

func CopyFile(from string, file1 string, file2 string) error {
	f, err := os.Open(from)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	d1, err := os.Create(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer d1.Close()

	d2, err := os.Create(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer d2.Close()

	for s.Scan() {
		d1.WriteString(s.Text() + "\n")
		d2.WriteString(s.Text() + "\n")

	}
	return nil
}

func FindWord(read string, word string) bool {
	f, err := os.Open(read)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		g := strings.Split(line, " ")
		for _, x := range g {
			x = NormTrimSuffix(x, ".", ",", "!", "?", "*")
			if x == word {
				return true
			}

		}
	}
	return false
}

func NormTrimSuffix(s string, suffix ...string) string {
	for _, x := range suffix {
		s = strings.TrimSuffix(s, x)
	}
	return s
}
