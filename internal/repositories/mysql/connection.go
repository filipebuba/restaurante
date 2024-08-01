package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	DB_HOST = "127.0.0.1"
	DB_PORT = 3306
	DB_NAME = "rest-db"
	DB_USER = "root"
	DB_PASS = "secret"
)

var db *sqlx.DB

func GetConnectionDB() (*sqlx.DB, error) {
	var err error

	if db == nil {
		db, err = sqlx.Connect("mysql", dbConnectionURL())
		if err != nil {
			fmt.Printf("########## DB ERROR: " + err.Error() + "##########")
			return nil, fmt.Errorf("### DB ERROR: %w", err)
		}

	}

	if err := migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *sqlx.DB) error {
	var clientSchema = `
	CREATE DATABASE IF NOT EXISTS restaurant;

	USE restaurant;

	CREATE TABLE IF NOT EXISTS clients (
		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
		nome varchar(20) DEFAULT NULL,
		email varchar(50) DEFAULT NULL,
		telefone tst,
		feedbacks text,
		orders text,
		PRIMARY KEY (id)
		UNIQUE KEY email (email)
	);`
	_, err := db.Exec(clientSchema)

	if err != nil {
		fmt.Printf("########## DB ERROR: " + err.Error() + "##########")
		return fmt.Errorf("### MIGRATION ERROR: %w", err)
	}

	return nil
}

func dbConnectionURL() string {
	return "root:root@tcp(localhost:3306)/restaurant?charset=utf8&parseTime=True"
}
