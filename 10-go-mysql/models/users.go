package models

import (
	"gomysql/db"
)

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(50) NOT NULL,
	email VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

func NewUser(username, password, email string) *User {
	return &User{Username: username, Password: password, Email: email}
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

func (user *User) insert() {
	query := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	result, _ := db.Exec(
		query,
		user.Username,
		user.Password,
		user.Email,
	)
	user.ID, _ = result.LastInsertId()
}

// Obtener todo el registro
func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

// Obtener un registro
func GetUser(id int64) *User {
	sql := "SELECT id, username, password, email FROM users WHERE id = ?"
	user := &User{}
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	}
	return user
}

// Actualizar un registro
func (user *User) update() {
	query := "UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?"
	db.Exec(query, user.Username, user.Password, user.Email, user.ID)
}

// guardar o editar registro
func (user *User) Save() {
	if user.ID == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// eliminar registro
func (user *User) Delete() {
	query := "DELETE FROM users WHERE id = ?"
	db.Exec(query, user.ID)
}
