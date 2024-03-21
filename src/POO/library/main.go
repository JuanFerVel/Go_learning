package main

import "library/animal"

func main() {
	// miPerro := animal.Perro{Nombre: "Lucas"}
	// miGato := animal.Gato{Nombre: "Mandarino"}

	// animal.HacerSonido(&miGato)
	// animal.HacerSonido(&miPerro)

	animales := []animal.Animal{
		&animal.Gato{Nombre: "Mandarino"},
		&animal.Perro{Nombre: "Max"},
		&animal.Perro{Nombre: "Lucas"},
		&animal.Gato{Nombre: "Dominic"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
