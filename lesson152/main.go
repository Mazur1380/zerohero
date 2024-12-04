package main

import (
	"log"
	"log/slog"
)

func main() {

	log.Println("Error new")
	slog.Info("Info 1", slog.String("User id", "1"))

}
