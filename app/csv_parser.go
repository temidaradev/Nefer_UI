package app

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to read input file: "+filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	csvReader.Comment = '#'

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for: "+filePath, err)
	}

	return records
}
