package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
)

// indexes referring to the columns of the customers table
const C_CPF = 0
const C_PRIVATE = 1
const C_INCOMPLETO = 2
const C_DATA_ULTIMA_COMPRA = 3
const C_TICKET_MEDIO = 4
const C_TICKET_ULTIMA_COMPRA = 5
const C_LOJA_MAIS_FREQUENTE = 6
const C_LOJA_ULTIMA_COMPRA = 7

func connectPostgresDB() *sql.DB {
	connstring := "user=postgres dbname=postgres password='postgres' host=go_db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("database created!")
	}
	return db
}

func createTableIntoPostgres(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS clients (id SERIAL PRIMARY KEY, cpf varchar(50), private integer, incompleto integer, data_ultima_compra date, ticket_medio varchar(50), ticket_ultima_compra varchar(50), loja_mais_frequente varchar(50), loja_ultima_compra varchar(50))")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("table created!")
	}
}

func insertIntoPostgres(db *sql.DB, colums_table map[int]string) {
	_, err := db.Exec("INSERT INTO  clients(cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_ultima_compra) VALUES($1,$2,$3,$4,$5,$6,$7,$8)", colums_table[0], colums_table[1], colums_table[2], colums_table[3], colums_table[4], colums_table[5], colums_table[6], colums_table[7])
	if err != nil {
		fmt.Println(err)
	}
}

func regexIsLetter(e string) bool {
	nameRegex, _ := regexp.Compile("[a-zA-Z]")
	return nameRegex.MatchString(e)
}

// func validator(e string) map[int]string {
// 	// nameRegex, _ := regexp.Compile("[a-zA-Z]")
// 	map_colums_clients_table := make(map[int]string)
// 	return map_colums_clients_table
// }

func main() {
	db := connectPostgresDB()
	createTableIntoPostgres(db)

	// Open the file
	file, err := os.Open("base_to_import/base_teste.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Read and print lines
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		} else {
			//	Map to index the values ​​of each line of the file, and insert them into the customers table
			map_colums_clients_table := make(map[int]string)
			map_columns_clients_index := 0

			if !regexIsLetter(line) {
				line_split := strings.SplitN(line, " ", -1)
				for line_index := 0; line_index < len(line_split); line_index++ {
					if line_split[line_index] != "" {
						map_colums_clients_table[map_columns_clients_index] = line_split[line_index]
						map_columns_clients_index++
					}
				}
				insertIntoPostgres(db, map_colums_clients_table)
			}
		}
	}
}
