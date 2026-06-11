package main

import (
	"database/sql"
	"strings"
	"unicode"
)
import "log"
import _ "modernc.org/sqlite"

func main() {

	db, err := sql.Open("sqlite", "phones.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) { _ = db.Close() }(db)

	query := `CREATE TABLE IF NOT EXISTS phone_numbers (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      number TEXT
  )`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return
	}

	numbers := []string{
		"1234567890",
		"0987654321",
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	for _, number := range numbers {
		_, err := db.Exec("INSERT INTO phone_numbers(number) VALUES (?)", number)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	type phoneRecord struct {
		id     int
		number string
	}

	rows, err := db.Query("SELECT id, number FROM phone_numbers")
	if err != nil {
		log.Fatal(err)
	}
	var records []phoneRecord
	for rows.Next() {
		var r phoneRecord
		err := rows.Scan(&r.id, &r.number)
		if err != nil {
			log.Fatal(err)
			return
		}
		records = append(records, r)
	}
	_ = rows.Close()

	uniqNumbers := map[string]struct{}{}
	for _, r := range records {

		normalizedNumber := normalizePhoneNumber(r.number)
		log.Printf("id: %d, normalized number: %s", r.id, normalizedNumber)

		if _, ok := uniqNumbers[normalizedNumber]; ok {
			log.Printf("duplicate number: %s", normalizedNumber)
			_, err := db.Exec("DELETE FROM phone_numbers WHERE id = ?", r.id)
			if err != nil {
				log.Fatal(err)
				return
			}
			continue
		}

		uniqNumbers[normalizedNumber] = struct{}{}

		_, err = db.Exec("UPDATE phone_numbers SET number=? WHERE id=?", normalizedNumber, r.id)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func normalizePhoneNumber(number string) string {
	b := strings.Builder{}
	for _, r := range number {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
