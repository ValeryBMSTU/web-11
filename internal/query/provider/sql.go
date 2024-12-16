package provider

import (
	"database/sql"
	"errors"
	"fmt"
)

func (p *Provider) FetchQuery(name string) (string, error) {
	var msg string

	err := p.conn.QueryRow("SELECT name_query FROM query WHERE name_query = $1", name).Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return "Hello, " + msg + "!", nil
}

func (p *Provider) CheckQueryExist(name string) (bool, error) {
	var msg string
	err := p.conn.QueryRow("SELECT name_query FROM query WHERE name_query = $1", name).Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *Provider) InsertQuery(name string) error {
	_, err := p.conn.Exec("INSERT INTO query (name_query) VALUES ($1)", name)
	if err != nil {
		return err
	}

	fmt.Println("Inserted successfully!")

	return nil
}
