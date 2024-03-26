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
		//Instancia del modelo Contact
		contact := models.Contact{}

		var valueEmail sql.NullString

		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "No registra correo"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Phone: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("-------------------------------------------------")
	}

}

//GetContactById traera un contacto  de la base de datos mediante un ID

func GetContactById(db *sql.DB, contactId int) {
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactId)

	contact := models.Contact{}
	var valueEmail sql.NullString

	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("Registro no encontrado con el id %d", contactId)

		} else {
			log.Fatal(err)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "No registra correo"
	}
	fmt.Println("LISTA DE UN CONTACTO")
	fmt.Println("-------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Phone: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("-------------------------------------------------")

}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contact(name, email, phone) VALUES(?, ?, ?)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto Registrado con Exito")
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact SET name = ?, email= ?, phone = ? WHERE id = ?"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("El registro se ha actualizado correctamente")
}

func DeleteContact(db *sql.DB, contactID int) {
	query := "DELETE FROM contact WHERE id = ?"
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto Eliminado con exito")
}
