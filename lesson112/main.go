package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now().UTC()
	fmt.Println(now)

	fmt.Println(now.Unix())

	t := time.Unix(1714835802, 0)
	fmt.Println(t.UTC())

	const shortForm = "2006/Jan/02-15:04:05"
	t, _ = time.Parse(shortForm, "2013/Feb/03-18:59:45")
	fmt.Println(t)

	loc, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		log.Fatal(err)
	}
	now = time.Now()
	fmt.Println(now)

	romeNow := now.In(loc)
	fmt.Println(romeNow)

	dur := "1m12s"

	d, _ := time.ParseDuration(dur)
	fmt.Println(d.Minutes())
}
