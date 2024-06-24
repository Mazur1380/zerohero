package main

import "fmt"

func main() {

	m := map[string][]string{
		"Alex": []string{"book1", "book2"},
		"Dima": []string{"book3", "book4"},
		"Vlad": []string{"book5", "book1"},
		"Djas": []string{"book3", "book2"},
	}
	z := Readers(m)
	fmt.Printf("Читателей у которых есть книга %d\n", z)
	v := Counter(m)
	fmt.Printf("Количество книг у читателей %d\n", v)
	c := CountUni(m)
	fmt.Printf("Колличетсво уникальных книг %d\n", c)
	a := CountUniMap(m)
	fmt.Printf("Колличетсво уникальных книг %d\n", a)
}

func Readers(m map[string][]string) int {
	c := 0
	for _, x := range m {
		if len(x) > 0 {
			c = c + 1
		}
	}
	return c
}

func Counter(m map[string][]string) int {
	c := 0
	for _, x := range m {
		c = c + len(x)
	}
	return c
}

func Exist(x []string, y string) bool {
	for _, v := range x {
		if v == y {
			return true
		}
	}
	return false
}

func CountUni(m map[string][]string) int {
	v := []string{}
	for _, x := range m {
		for _, g := range x {
			if Exist(v, g) {
				continue
			}
			v = append(v, g)
		}
	}
	return len(v)
}

func CountUniMap(m map[string][]string) int {
	c := map[string]struct{}{}
	for _, d := range m {
		for _, r := range d {
			if _, ok := c[r]; !ok {
				c[r] = struct{}{}
			}
		}
	}
	return len(c)

}
