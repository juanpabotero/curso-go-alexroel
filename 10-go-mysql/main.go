package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()

	// db.CreateTable(models.UserSchema, "users")
	// db.TruncateTable("users")

	// models.CreateUser("Aleja", "123", "aleja@gmail.com")

	// user := models.GetUser(4)

	// user.Username = "Olga Ber"
	// user.Save()

	// user.Delete()
	// db.TruncateTable("users")
	users := models.ListUsers()
	fmt.Println(users)

	db.Close()
}
