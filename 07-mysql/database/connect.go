package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	// cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// contiene los datos de la base de datos
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Conectar a la base de datos
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	// verificar conexion
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión exitosa")

	return db, nil
}
