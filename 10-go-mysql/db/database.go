package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func Connect() {
	// cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		return
	}

	// contiene los datos de la base de datos
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// conectar a la base de datos
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	db = connection
}

func Close() {
	// cerrar conexión a la base de datos
	db.Close()
}

func Ping() {
	// verificar conexión a la base de datos
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Ping to database")
}

func TableExists(name string) bool {
	// verificar si la tabla existe en la base de datos
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", name)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	// cerrar filas al final de la función
	defer rows.Close()
	// si rows.Next() retorna true, la tabla existe
	// Next avanza al siguiente registro, si puede avanzar, significa que la tabla existe
	return rows.Next()
}

func CreateTable(schema, name string) {
	// verificar si la tabla existe
	if TableExists(name) {
		return
	}
	// crear tabla en la base de datos
	_, err := db.Exec(schema)
	if err != nil {
		panic(err)
	}
}

// reiniciar el registro de una tabla
func TruncateTable(name string) {
	// verificar si la tabla existe
	if !TableExists(name) {
		return
	}
	// reiniciar el registro de una tabla
	query := fmt.Sprintf("TRUNCATE TABLE %s", name)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// polimorfismo de exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	// ejecutar una consulta en la base de datos
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Polimorfismo a Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	// ejecutar una consulta en la base de datos
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err.Error())
	}
	return rows, err
}
