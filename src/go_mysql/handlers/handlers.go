package handlers

import (
	"database/sql"
	"fmt"
	"go_mysql/models"
	"log"
)

// Lsitar contactos desde la base de datos
func ListContacts(db *sql.DB) {
	//Consulta SQL
	query := "SELECT * FROM contact"

	//Ejecutar la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("LISTA DE CONTACTOS")
	fmt.Println("-------------------------------------------------")
	for rows.Next() {
		//Instancia del modelo Contyact
		contact := models.Contact{}
		err := rows.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Phone: %s\n",
			contact.Id,
			contact.Name,
			contact.Email,
			contact.Phone)
	}

}
