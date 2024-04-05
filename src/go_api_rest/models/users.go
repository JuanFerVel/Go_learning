package models

import "go_api_rest/db"

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`

	// Id       int64  `xml:"id"`
	// UserName string `xml:"username"`
	// Password string `xml:"password"`
	// Email    string `xml:"email"`

	// Id       int64  `yaml:"id"`
	// UserName string `yaml:"username"`
	// Password string `yaml:"password"`
	// Email    string `yaml:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

// Construir Usuario
func NewUser(username, password, email string) *User {
	user := &User{UserName: username, Password: password, Email: email}
	return user
}

// Crear Usuario e inserta BD
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// Insertar Registro
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.UserName, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

// Listar Registros
func ListUsers() (Users, error) {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, err := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users, err
}

// Listar un unico registro
func GetUser(id int) (*User, error) {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	if rows, err := db.Query(sql, id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		}
		return user, nil
	}
}

// Editar Usuario
func (user *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.UserName, user.Password, user.Email, user.Id)
}

// Guardar o editar Registro
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// Eliminar un registro
func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}
