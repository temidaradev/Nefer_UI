package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() *sql.DB {
	if _, err := os.Stat("./data/values.db"); err != nil {
		db, err := sql.Open("sqlite3", "./data/values.db")
		if err != nil {
			fmt.Println(err)
		}

		return db
	} else {
		db, err := sql.Open("sqlite3", "./data/values.db")
		if err != nil {
			fmt.Println(err)
		}
		CreateTable(db)
		return db
	}
}

func CreateTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS "values" (
		Time FLOAT,
		Speed INT,
		Temp FLOAT,
		Voltage INT,
		Watt FLOAT
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("aaaaaaaaaaaaaa", err)
	}
}
