package postgres

import (
	"database/sql"
	"fmt"
	constants "goimporttodb/utils"
)

// Function do connect on DB
func ConnectPostgresDB() *sql.DB {
	connstring := "user=postgres dbname=postgres password='postgres' host=go_db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database created!")
	}
	return db
}

// Function to create talble
func CreateTableIntoPostgres(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS clients (id SERIAL PRIMARY KEY, cpf varchar(50), private integer, incompleto integer, data_ultima_compra date, ticket_medio real, ticket_ultima_compra real, loja_mais_frequente varchar(50), loja_ultima_compra varchar(50))")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Table created!")
	}
}

// Function to insert data on DB
func InsertIntoPostgres(db *sql.DB, colums_table map[int]string) {
	_, err := db.Exec("INSERT INTO clients(cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_ultima_compra) VALUES($1,$2,$3,$4,$5,$6,$7,$8)",
		colums_table[constants.COLUMN_CPF],
		colums_table[constants.COLUMN_PRIVATE],
		colums_table[constants.COLUMN_INCOMPLETO],
		colums_table[constants.COLUMN_DATA_ULTIMA_COMPRA],
		colums_table[constants.COLUMN_TICKET_MEDIO],
		colums_table[constants.COLUMN_TICKET_ULTIMA_COMPRA],
		colums_table[constants.COLUMN_LOJA_MAIS_FREQUENTE],
		colums_table[constants.COLUMN_LOJA_ULTIMA_COMPRA])
	if err != nil {
		fmt.Println(err)
	}
}
