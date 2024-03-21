/*Ejercicio
Ejercicio: Calcule e imprima el área y el perímetro del triángulo
Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.
El programa debe:

Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.
Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.
Calcular el área del triángulo usando la fórmula base x altura / 2.
Calcular el perímetro del triángulo sumando los lados.
Imprimir el área y el perímetro del triángulo con dos decimales de precisión.
El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro. También debe usar constantes para representar el número de decimales de precisión que se desean en la salida.
Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.

Ejemplo de entrada y salida:

Ingrese lado 1: 3.5
Ingrese lado 2: 4.2

Área: 7.35
Perímetro: 12.20*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var h, b, a, c, p float64
	//h -> Altura
	//b -> Base
	//a -> Área
	//c -> Hipotenusa
	//p -> Perimetro

	//Solicitud de Datos
	fmt.Println("------------------- Ejercicio -------------------")
	fmt.Print("Ingrese la altura del triangulo rectangulo: ")
	fmt.Scanln(&h)
	fmt.Print("Ingrese la base del triangulo rectangulo: ")
	fmt.Scanln(&b)

	//Calculo de Hipotenusa
	d := math.Pow(h, 2) + math.Pow(b, 2)
	c = math.Sqrt(d)

	//Calculo de área
	a = (b * h) / 2

	//Calculo del perimetro
	p = h + b + c

	// x := fmt.Sprintf("%.2f", c)
	y := fmt.Sprintf("%.2f", p)
	z := fmt.Sprintf("%.2f", a)

	// fmt.Printf("Segun los catetos a: %v y b: %v, el valor de la hipotenusa es: %v \n", h, b, x)
	fmt.Printf("El área del triangulo rectangulo es %v \n", z)
	fmt.Printf("El Perimetro del triangulo rectangulo es %v \n", y)
}
