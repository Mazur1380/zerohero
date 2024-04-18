package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	dbUrl := "postgres://postgres:pass@localhost:5432/postgres"
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	// insert(ctx, conn, "dima", "444", "Dima", 24)
	// if err != nil {
	// log.Fatal(err)
	// }
	id, login, password, name, age, err := get(ctx, conn, "mazur123")
	if errors.Is(err, pgx.ErrNoRows) {
		fmt.Println("Пользователь не найден")
		return
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Нашли аккаунт:", id, login, password, name, age)
}

func insert(ctx context.Context, conn *pgx.Conn, login string, password string, name string, age int) error {
	const query = "INSERT INTO accounts (login,password,name,age)VALUES($1,$2,$3,$4);"
	_, err := conn.Exec(ctx, query, login, password, name, age)
	if err != nil {
		return err
	}
	return nil

}
func get(ctx context.Context, conn *pgx.Conn, login string) (int, string, string, string, int, error) {
	const query = "SELECT id,login,password,name,age FROM accounts WHERE login = $1;"
	row := conn.QueryRow(ctx, query, login)

	var id int
	var accountLogin, password, name string
	var age int

	err := row.Scan(&id, &accountLogin, &password, &name, &age)
	if err != nil {
		return 0, "", "", "", 0, err
	}
	return id, accountLogin, password, name, age, nil
}
