package main

import (
	"fmt"
	"go_Web/db"
	"go_Web/models"
)

func main() {
	db.Connect()
	// db.Ping()
	// fmt.Println(db.ExistTable("users"))
	// db.CreateTale(models.UserSchema, "users")
	db.TruncateTable("users")

	user := models.CreateUser("Lucas Martinez", "Lumas123", "lumartines@gmail.com")
	fmt.Println(user)

	// users := models.ListUsers()
	// fmt.Println(users)

	// user := models.GetUser(2)
	// user.UserName = "Juan"
	// user.Password = "juan789"
	// user.Email = "juan@gmail.com"
	// user.Save()
	// users := models.ListUsers()
	// fmt.Println(users)
	// user.Delete()
	users := models.ListUsers()
	fmt.Println(users)
	db.Close()
}
