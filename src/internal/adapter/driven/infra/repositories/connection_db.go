package infra

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewConnection() *pgx.Conn {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBname)

	conn, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		log.Printf("")
		panic(1)
	}

	return conn
}
