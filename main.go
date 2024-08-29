package main

import (
	"analyzer/database"
	"analyzer/processing"
	"fmt"
	"log"
	"os"
)

func main() {
	
	files, err := processing.Unzip("C:/Users/My PC/Downloads/Dataset.csv.zip", "unzipped_files")
	if err != nil {
		log.Fatalf("Error unzipping files: %v", err)
		os.Exit(1)
	}
	fmt.Println("Unzipped files:", files)

	for _, file := range files {
		records, err := processing.ReadCSV(file)
		if err != nil {
			log.Printf("Error reading CSV file %s: %v", file, err)
			continue
		}

		cleanData := processing.CleanData(records)

		db, err := database.ConnectToDB()
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}
		defer db.Close()

		err = database.InsertData(db, cleanData)
		if err != nil {
			log.Printf("Error inserting data into the database for file %s: %v", file, err)
			continue
		}
	}

}
