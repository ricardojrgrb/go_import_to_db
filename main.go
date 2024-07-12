package main

import (
	"bufio"
	"fmt"
	"goimporttodb/postgres"
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
	db := postgres.ConnectPostgresDB()
	postgres.CreateTableIntoPostgres(db)

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
				postgres.InsertIntoPostgres(db, map_colums_clients_table)
			}
		}
	}
}
