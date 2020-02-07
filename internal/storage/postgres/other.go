package postgres

import (
	"database/sql"
	"errors"
)

func DatabaseConnectionCheck(connections ...*sql.DB) error {
	for _, conn := range connections {
		if conn == nil {
			return errors.New("connection_nil")
		}
		if _, err := conn.Exec("SELECT 1 + 1;"); err != nil {
			return errors.New("connection_dead")
		}
	}
	return nil
}
