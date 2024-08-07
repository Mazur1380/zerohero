package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	err := CopyFile("in.txt", "out.txt")
	if err != nil {
		log.Fatal(err)
	}

	counter, err := CountLetter("in.txt", "into.txt", "b")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counter)
}

func CopyFile(from string, to string) error {
	f, err := os.Open(from)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	d, err := os.Create(to)
	if err != nil {
		return err
	}
	defer d.Close()

	for s.Scan() {
		d.WriteString(s.Text() + "\n")
	}
	return nil

}

func CountLetter(from string, to string, letter string) (int, error) {
	f, err := os.Open(from)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	d, err := os.Create(to)
	if err != nil {
		return 0, err
	}
	defer d.Close()

	counter := 0

	for s.Scan() {
		text := s.Text()
		for _, x := range text {
			if string(x) == letter {
				counter += 1
			}
		}
	}
	e := fmt.Sprintf("Буква %s встречается %d  раз", letter, counter)

	d.WriteString(e)

	return counter, nil

}
