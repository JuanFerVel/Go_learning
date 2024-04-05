package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// const url = "root:FerJuanVel#14@tcp(localhost:3000)/go_web" //URL PERSONAL COMPUTER
const url = "juanfervel:password@tcp(localhost:3306)/go_web" // URL EMPRESARIAL COMPUTER

var db *sql.DB

// Conecta a la Base de Datos
func Connect() {
	con, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion Exitosa")
	db = con
}

// Cierra la conexion de la base de datos
func Close() {
	db.Close()
}

// Verifica si al conexion sigue abierta
func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

// verifica si una tabla existe en la BD
func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

// Crea una tabla inexistente en la base de datos
func CreateTale(schema string, tableName string) {
	if !ExistTable(tableName) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Reiniciar los Datos de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

// Polimorfismo EXEC
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// Polimorfismo QUERY
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
