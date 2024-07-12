package postgres

import (
	"database/sql"
	"fmt"
)

func ConnectPostgresDB() *sql.DB {
	connstring := "user=postgres dbname=postgres password='postgres' host=go_db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("database created!")
	}
	return db
}

func CreateTableIntoPostgres(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS clients (id SERIAL PRIMARY KEY, cpf varchar(50), private integer, incompleto integer, data_ultima_compra date, ticket_medio varchar(50), ticket_ultima_compra varchar(50), loja_mais_frequente varchar(50), loja_ultima_compra varchar(50))")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("table created!")
	}
}

func InsertIntoPostgres(db *sql.DB, colums_table map[int]string) {
	_, err := db.Exec("INSERT INTO  clients(cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_ultima_compra) VALUES($1,$2,$3,$4,$5,$6,$7,$8)", colums_table[0], colums_table[1], colums_table[2], colums_table[3], colums_table[4], colums_table[5], colums_table[6], colums_table[7])
	if err != nil {
		fmt.Println(err)
	}
}
