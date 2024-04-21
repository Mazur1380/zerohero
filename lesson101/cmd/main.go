package main

import (
	"context"
	"fmt"
	"log"

	"github.com/andreipimenov/zerohero/lesson101/internal/store"
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

	st := store.Store{
		Conn: conn,
	}

	acc, err := st.GetAccount(ctx, "elvin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(acc)

	// acc := model.Account{
	// 	Login:    "elvin",
	// 	Password: "333",
	// 	Name:     "Elvin",
	// 	Age:      25,
	// }
	// err = st.InsertAccount(ctx, acc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
