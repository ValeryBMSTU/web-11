package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) CreateUser(username, password string) error {
	_, err := p.conn.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	return err
}

// func (p *Provider) SelectRandomHello() (string, error) {
// 	var msg string

// 	// Получаем одно сообщение из таблицы hello, отсортированной в случайном порядке
// 	err := p.conn.QueryRow("SELECT message FROM hello ORDER BY RANDOM() LIMIT 1").Scan(&msg)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return "", nil
// 		}
// 		return "", err
// 	}

// 	return msg, nil
// }

func (p *Provider) CheckUserByUsername(username string) (bool, error) {
	err := p.conn.QueryRow("SELECT (username) FROM users WHERE username = $1", username).Scan(&username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *Provider) CheckPassword(username, password string) (bool, error) {
	var password_db string
	err := p.conn.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&password_db)
	if err != nil {
		return false, err
	}
	if password == password_db {
		return true, nil
	}
	return false, nil
}

// func (p *Provider) CheckHelloExitByMsg(msg string) (bool, error) {
// 	// Получаем одно сообщение из таблицы hello
// 	err := p.conn.QueryRow("SELECT message FROM hello WHERE message = $1 LIMIT 1", msg).Scan(&msg)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return false, nil
// 		}
// 		return false, err
// 	}

// 	return true, nil
// }

// func (p *Provider) InsertHello(msg string) error {
// 	_, err := p.conn.Exec("INSERT INTO hello (message) VALUES ($1)", msg)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
