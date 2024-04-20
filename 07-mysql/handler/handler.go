package handler

import (
	"07-mysql/model"
	"database/sql"
	"fmt"
	"log"
)

// listar contactos
func ListContacts(db *sql.DB) {
	query := "SELECT * FROM contact"
	// ejecutar query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	// cerrar filas
	// se debe cerrar porque rows es un recurso que se debe liberar
	defer rows.Close()
	// recorrer filas
	for rows.Next() {
		// instancia del modelo
		contact := model.Contact{}

		var valueEmail sql.NullString

		// escanear filas, lee los valores de las columnas y los almacena en las variables
		// se debe pasar el mismo numero de argumentos que las columnas seleccionadas
		err := rows.Scan(&contact.ID, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		// validar si el email es nulo
		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "No tiene email"
		}

		// imprimir datos
		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	}
	// verificar si hubo un error en rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// obtener un contacto
func GetContactById(db *sql.DB, id int) {
	query := "SELECT * FROM contact WHERE id = ?"
	// ejecutar query
	row := db.QueryRow(query, id)

	// instancia del modelo
	contact := model.Contact{}

	var valueEmail sql.NullString

	// escanear fila
	err := row.Scan(&contact.ID, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No existe un contacto con el ID %d", id)
		}
		log.Fatal(err)
	}

	// validar si el email es nulo
	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "No tiene email"
	}

	// imprimir datos
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
}

// crear un contacto
func CreateContact(db *sql.DB, contact model.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"
	// ejecutar query
	result, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	// obtener el ID del contacto creado
	id, _ := result.LastInsertId()
	// imprimir mensaje
	fmt.Printf("Contacto creado con el ID %d\n", id)
}

// editar contacto
func UpdateContact(db *sql.DB, id int, contact model.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"
	// ejecutar query
	result, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, id)
	if err != nil {
		log.Fatal(err)
	}
	// obtener el número de filas afectadas
	rows, _ := result.RowsAffected()
	// imprimir mensaje
	fmt.Printf("Contacto actualizado con %d filas afectadas\n", rows)
}

// eliminar un contacto
func DeleteContact(db *sql.DB, id int) {
	query := "DELETE FROM contact WHERE id = ?"
	// ejecutar query
	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	// obtener el número de filas afectadas
	rows, _ := result.RowsAffected()
	// imprimir mensaje
	fmt.Printf("Contacto eliminado con %d filas afectadas\n", rows)
}
