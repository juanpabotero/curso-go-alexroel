package models

import "gorm/db"

type User struct {
	Id       int64  `json:"id"`
	Usename  string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

func MigrarUser() {
	// migrar el esquema de la base de datos
	db.Database.AutoMigrate(User{})
}