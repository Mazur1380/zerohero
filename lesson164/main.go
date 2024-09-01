package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Выполняем команду `ls`
	cmd := exec.Command("docker", "run", "-d", "-p=80:80", "nginx")

	// Создаем байтовый буфер для захвата вывода команды
	var out bytes.Buffer
	cmd.Stdout = &out

	// Выполняем команду и проверяем на наличие ошибок
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// Выводим результат на экран
	fmt.Println(out.String())
}
