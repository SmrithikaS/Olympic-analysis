package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	connStr := "user:abcd@tcp(127.0.0.1:3606)/mydb"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}
	return db, nil
}

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS athletes (
		ID INT,
		Name VARCHAR(255),
		Sex VARCHAR(255),
		Age INT,
		Height INT,
		Weight INT,
		Team VARCHAR(255),
		NOC VARCHAR(255),
		Games VARCHAR(255),
		Year INT,
		Season VARCHAR(255),
		City VARCHAR(255),
		Sport VARCHAR(255),
		Event VARCHAR(255),
		Medal VARCHAR(255)
	);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table")
	}
	return nil
}

func InsertData(db *sql.DB, data [][]string) error {
	err := CreateTable(db)
	if err != nil {
		return err
	}

	insertQuery := `
	INSERT INTO athletes (ID,Name,Sex, Age, Height, Weight, Team, NOC, Games, Year, Season, City, Sport, Event, Medal)
	VALUES ("%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v")`

	for _, record := range data[1:] {
		query := fmt.Sprintf(insertQuery, record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14])
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("error inserting data")
		}
	}

	return nil
}
