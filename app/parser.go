package app

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Value struct {
	Time    float64 `db:"time"`
	Speed   int     `db:"speed"`
	Temp    float64 `db:"temp"`
	Voltage int     `db:"voltage"`
	Watt    float64 `db:"watt"`
}

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

	if err == io.EOF {
		fmt.Println("niga")
	}

	return records[1:]

}

func ParseCsv(filepath string) ([][]string, []string) {

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var records [][]string
	for {

		record, err := reader.Read()
		// if we've reached the end of the file, break
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		records = append(records, record)
	}

	return records[1:], records[0]

}

func stringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func stringToInt(s string) int {
	s = strings.TrimSuffix(s, ".0")
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseData(records [][]string) []Value {
	var values []Value
	for _, record := range records {
		value := Value{
			Time:    stringToFloat(record[0]),
			Speed:   stringToInt(record[1]),
			Temp:    stringToFloat(record[2]),
			Voltage: stringToInt(record[3]),
			Watt:    stringToFloat(record[4]),
		}
		values = append(values, value)
	}
	return values
}

func FindAll(db *sql.DB) []Value {
	// Escape the table name with double quotes
	rows, err := db.Query(`SELECT * FROM "values"`)
	if err != nil {
		log.Fatal("Query error:", err)
	}
	defer rows.Close()

	var values []Value

	for rows.Next() {
		var value Value
		// Ensure fields are exported (capitalized) in your Value struct
		err := rows.Scan(&value.Time, &value.Speed, &value.Temp, &value.Voltage, &value.Watt)
		if err != nil {
			log.Fatal("Scan error:", err)
		}
		values = append(values, value)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Rows error:", err)
	}

	return values
}

func InsertData(db *sql.DB, values []Value) {
	for _, value := range values {
		// the ? allows us to send the values during execution
		sqlStmt := `
  		INSERT INTO "values"(time,speed,temp,voltage,watt) VALUES(?,?,?,?,?)
  		`

		_, err := db.Exec(sqlStmt, &value.Time, &value.Speed, &value.Temp, &value.Voltage, &value.Watt)
		if err != nil {
			fmt.Println("eeeeeeeeeeeee", err)
		}
	}
}
