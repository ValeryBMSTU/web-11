package provider

import (
	"database/sql"
	"fmt"
	"log"
)

type Provider struct {
	conn *sql.DB
}

func NewProvider(host string, port int, user, dbName string) *Provider {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbName)

	// Создание соединения с сервером postgres
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{conn: conn}
}
