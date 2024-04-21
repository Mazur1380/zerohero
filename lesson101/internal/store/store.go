package store

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/andreipimenov/zerohero/lesson101/internal/model"
)

type Store struct {
	Conn *pgx.Conn
}

func (s *Store) InsertAccount(ctx context.Context, acc model.Account) error {
	const query = "INSERT INTO accounts (login,password,name,age)VALUES($1,$2,$3,$4);"

	_, err := s.Conn.Exec(ctx, query, acc.Login, acc.Password, acc.Name, acc.Age)
	if err != nil {
		return err
	}
	return nil

}
func (s *Store) GetAccount(ctx context.Context, login string) (model.Account, error) {
	const query = "SELECT id,login,password,name,age FROM accounts WHERE login = $1;"

	var acc model.Account

	err := s.Conn.QueryRow(ctx, query, login).Scan(&acc.ID, &acc.Login, &acc.Password, &acc.Name, &acc.Age)
	if err != nil {
		return model.Account{}, err
	}
	return acc, nil
}
