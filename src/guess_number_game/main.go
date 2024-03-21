package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(time.Now().UnixNano()) //Por si no genera el número Añeatorio
	fmt.Println(rand.Intn(100))
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100) + 1
	var numIngresado, intentos, restIntentos int
	const maxIntentos = 10

	for intentos = 1; intentos <= maxIntentos; intentos++ {
		restIntentos = maxIntentos - intentos + 1
		fmt.Printf("Ingresa un número (Intentos restantes %d):", restIntentos)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Felicitaciones Adivinaste el número")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a adivinar es mayor ")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a adivinar es menor ")
		}
	}
	fmt.Printf("Se acabaron los intentos gracias por jugar, el número era %d \n", numAleatorio)
	jugarNuevamente()
}
func jugarNuevamente() {
	var opcion string
	fmt.Println("Quieres Jugar de nuevo? (S/N)")
	fmt.Scanln(&opcion)

	switch opcion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Opción Invalida")
		jugarNuevamente()
	}
}
