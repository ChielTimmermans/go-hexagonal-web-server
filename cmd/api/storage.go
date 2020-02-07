package main

import (
	"database/sql"
	"fmt"
	"go-hexagonal/internal"
	"go-hexagonal/internal/domain/user"
	"go-hexagonal/internal/storage/postgres"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	user user.Storager
}

func initStorage(csPostgres *ConfigStoragePostgres, databaseStorageMode string) (*Storage, error) {
	s := &Storage{}

	if err := initDatabase(csPostgres, s, databaseStorageMode); err != nil {
		return nil, err
	}
	return s, nil
}

func initDatabase(config *ConfigStoragePostgres, s *Storage, databaseMode string) (err error) {
	log.Printf("Connecting to %s", databaseMode)
	switch databaseMode {
	case "postgres":
		var db *sql.DB
		if db, err = NewPostgresDatabaseConn(config.Hostname, config.User, config.Password, config.Database, config.Port); err != nil {
			return err
		}
		if s.user, err = postgres.NewUserStorage(db); err != nil {
			return err
		}
	default:
		return fmt.Errorf("storage mode unknown. \t possible modes: %s\t given mode: %s", possibleModes([]string{"vitess"}), databaseMode)
	}
	return nil
}

func possibleModes(possibleDBModes []string) string {
	var data string
	for _, v := range possibleDBModes {
		data = internal.Concat(data, " ['", v, "']")
	}
	data = internal.Concat(data, " ")
	return data
}

func NewPostgresDatabaseConn(hostname, user, password, database string, port int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, port, user, password, database)

	log.Println("Connecting to DB...")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to DB")

	return db, nil
}
