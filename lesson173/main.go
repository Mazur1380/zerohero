package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string = "104"
	var b int = 35

	x, _ := strconv.Atoi(a)
	y := strconv.Itoa(b)

	fmt.Println(x, y)
}

//объявите две переменные, первая - строка со значением 104, вторая - целое число со значением 35;
//приведите строку к целому числу, а целое число - к строке;
