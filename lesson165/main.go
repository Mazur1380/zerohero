package main

import (
	"fmt"
	"regexp"
)

func isValidEmail(email string) bool {
	// Регулярное выражение для проверки email
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func main() {
	var email string

	fmt.Print("Введите email: ")
	fmt.Scan(&email)

	if isValidEmail(email) {
		fmt.Println("Email валиден.")
	} else {
		fmt.Println("Email невалиден.")
	}
}
