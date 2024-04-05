package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:FerJuanVel#14@tcp(localhost:3000)/go_web"

var db *sql.DB

func Connect() {
	con, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion Exitosa")
	db = con
}

func Close() {
	db.Close()
}

func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

func CreateTale(schema string, tableName string) {
	if !ExistTable(tableName) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}
