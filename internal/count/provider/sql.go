package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) FetchCount() (int, error) {
	var msg int

	err := p.conn.QueryRow("SELECT number FROM counter").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return msg, nil
}

func (p *Provider) CheckCountExist() (bool, error) {
	msg := 1
	err := p.conn.QueryRow("SELECT number FROM counter WHERE id_number = $1", msg).Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *Provider) UpdateCount(count int) error {
	_, err := p.conn.Exec("UPDATE counter SET number = number + $1 WHERE id_number = 1", count)
	if err != nil {
		return err
	}

	return nil
}
