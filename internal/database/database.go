package database

import (
	"database/sql"
	"fmt"
)

type DbConn struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func SetupDatabase(cfg *DbConn) (*sql.DB, error) {
	dbConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
