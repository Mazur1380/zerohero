package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка получения текущей директории:", err)
		return
	}

	fmt.Println("Содержимое папки:", dir)

	if err := OpenFile(dir, "-"); err != nil {
		fmt.Println("Ошибка при выводе директории:", err)
	}
}

func OpenFile(catalog string, prefix string) error {
	d, err := os.ReadDir(catalog)
	if err != nil {
		log.Fatal(err)
	}

	for _, x := range d {
		fmt.Println(prefix, x.Name())
		if x.IsDir() {
			newCatalog := filepath.Join(catalog, x.Name())
			newPrefix := prefix + "-"
			if err := OpenFile(newCatalog, newPrefix); err != nil {
				return err
			}
		}
	}
	return nil
}
