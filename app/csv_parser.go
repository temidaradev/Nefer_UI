package app

import (
	"encoding/csv"
	"fmt"
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

	fields := make(map[string]int)
	for i, name := range records[0] {
		fields[name] = i
	}

	for _, record := range records[1:] {
		x := record[fields["zaman_ms"]]
		fmt.Println(x)
	}
	for _, record := range records[0:] {
		x := record[fields["hiz_kmh"]]
		fmt.Println(x)
	}

	return records
}
