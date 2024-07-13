package main

import (
	"bufio"
	"fmt"
	"goimporttodb/normalization"
	"goimporttodb/postgres"
	constants "goimporttodb/utils"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	// Connect on DB
	db := postgres.ConnectPostgresDB()
	// Create table
	postgres.CreateTableIntoPostgres(db)

	// Open the file
	file, err := os.Open(constants.PATH_OF_FILE_TO_SPLIT_LINES)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)
	fmt.Println("File processing started!")
	// Read file
	for {
		// Read line per line
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		} else {
			//	Map to index the values ​​of each line of the file, and insert them into the customers table
			map_colums_clients_table := make(map[int]string)
			// index to populate map table
			index_columns_clients := 0

			line_split := strings.SplitN(line, " ", -1)
			for line_index := 0; line_index < len(line_split); line_index++ {
				// Checks index by index of the read line and stores indexes other than white space in the map table
				if line_split[line_index] != "" {
					map_colums_clients_table[index_columns_clients] = line_split[line_index]
					index_columns_clients++
				}
			}

			// Map to receive normalized data of line
			map_colums_clients_table_normalized := make(map[int]string)
			map_colums_clients_table_normalized = normalization.NormalizationLine(map_colums_clients_table)
			// If not a invalid line, persists in the database
			if len(map_colums_clients_table_normalized) != 0 {
				postgres.InsertIntoPostgres(db, map_colums_clients_table_normalized)
			}
		}
	}
	fmt.Println("Finished file processing!")
}
