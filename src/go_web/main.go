package main

import (
	"fmt"
	"go_Web/db"
	"go_Web/models"
)

func main() {
	db.Connect()
	db.Ping()
	fmt.Println(db.ExistTable("users"))
	db.CreateTale(models.UserSchema, "users")
	db.Close()
}
