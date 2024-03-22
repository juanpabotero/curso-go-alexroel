package main

import (
	"07-mysql/database"
	"07-mysql/handler"
	"07-mysql/model"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	// _ significa que se usa el paquete pero no se llama directamente
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// establecer conexion a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	// cerrar la conexion a la base de datos
	// defer se usa para ejecutar una función despues de que la función que lo contiene termine
	defer db.Close()

	for {
		// menu
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		// leer opción
		fmt.Print("Opción: ")
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			// listar contactos
			handler.ListContacts(db)
		case 2:
			// obtener un contacto
			fmt.Print("ID: ")
			var id int
			fmt.Scanln(&id)
			handler.GetContactById(db, id)
		case 3:
			// crear un contacto
			newContact := inpurContactDetails(3)
			handler.CreateContact(db, newContact)
		case 4:
			// actualizar un contacto
			updatedContact := inpurContactDetails(4)
			handler.UpdateContact(db, updatedContact.ID, updatedContact)
		case 5:
			// eliminar un contacto
			fmt.Print("ID: ")
			var id int
			fmt.Scanln(&id)
			handler.DeleteContact(db, id)
		case 6:
			// salir
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}

// obtener los datos el contacto
func inpurContactDetails(option int) model.Contact {
	// leer la entrada del ususario usando buffio
	// bufio.NewReader(os.Stdin) crea un nuevo objeto bufio.Reader que lee de os.Stdin
	reader := bufio.NewReader(os.Stdin)

	// crear un nuevo contacto
	var contact model.Contact

	if option == 4 {
		fmt.Print("ID: ")
		var id int
		fmt.Scanln(&id)
		contact.ID = id
	}

	fmt.Print("Nombre: ")
	// bufio.Reader.ReadString('\n') lee hasta que encuentra un salto de línea
	name, _ := reader.ReadString('\n')
	// strings.TrimSpace() elimina los espacios en blanco al inicio y al final de la cadena
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Phone: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
