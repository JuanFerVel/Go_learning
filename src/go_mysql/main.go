package main

import (
	"bufio"
	"fmt"
	"go_mysql/database"
	"go_mysql/handlers"
	"go_mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" // la "_" es porque se esta usando la libreria de forma indirecta
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		var option int
		fmt.Print("Menú: \n",
			"1. Listar Contactos \n",
			"2. Obtener un Contacto por id \n",
			"3. Crear un nuevo contacto \n",
			"4. Actualizar un contacto \n",
			"5.	Eliminar un contacto \n",
			"6. Salir \n",
			"Seleccione una opción: ")
		fmt.Scanln(&option)
		switch option {
		case 1:
			handlers.ListContacts(db)

		case 2:
			var idSeacrh int
			fmt.Print("\n Digite el ID del registro que desea Buscar: ")
			fmt.Scanln(&idSeacrh)
			handlers.GetContactById(db, idSeacrh)

		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)

		case 4:
			newContact := inputContactDetails(option)
			handlers.UpdateContact(db, newContact)
			handlers.ListContacts(db)

		case 5:
			var idSeacrh int
			fmt.Print("\n Digite el ID del registro que desea Eliminar: ")
			fmt.Scanln(&idSeacrh)
			handlers.DeleteContact(db, idSeacrh)
		case 6:
			return
		default:
			fmt.Println("Opcion no valida, intente de nuevo")
		}
	}

	// handlers.ListContacts(db)

	// handlers.GetContactById(db, 4)

	// newContact := models.Contact{
	// 	Id:    4,
	// 	Name:  "Juan Desalcachofa",
	// 	Email: "juan@gmail.com",
	// 	Phone: "1234567890",
	// }
	// handlers.UpdateContact(db, newContact)
	// handlers.CreateContact(db, newContact)
	// handlers.DeleteContact(db, 5)

}

func inputContactDetails(option int) models.Contact {
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		fmt.Print("\n Digite el ID del contacto que desea actualizar: ")
		fmt.Scanln(&contact.Id)
	}

	fmt.Print("\n Digite el Nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("\n Digite el Email del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("\n Digite el Telefono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
