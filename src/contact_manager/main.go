package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de contactos
type Contact struct {
	Name  string `jeson:"name"`
	Email string `jeson:"email"`
	Phone string `jeson:"phone"`
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("Contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func loadContactFronFile(contacts *[]Contact) error {
	file, err := os.Open("Contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return err
}

func main() {
	//Slice de contactos
	var contacts []Contact

	//Cargar contactos existentes desde el archivo
	err := loadContactFronFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos")
	}

	//Crear Instancia de bufio
	reader := bufio.NewReader(os.Stdin)
	for {
		var option int
		fmt.Print("==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar Contacto\n",
			"2. Mostrar Todos los Contactos\n",
			"3. Salir\n",
			"Escoge una opción: ")
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}
		case 2:
			fmt.Println("=================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Phone: %s", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=================================")
		case 3:
			return
		default:
			fmt.Println("Opción Invalida")
		}
	}

}
