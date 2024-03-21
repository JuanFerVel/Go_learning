/*Ejercicio
Ejercicio: Calcule e imprima el área y el perímetro del triángulo
Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.
El programa delado2e:

Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.
Calcular la lado1ipotenusa del triángulo usando el teorema de Pitágoras.
Calcular el área del triángulo usando la fórmula lado2ase x altura / 2.
Calcular el perímetro del triángulo sumando los lados.
Imprimir el área y el perímetro del triángulo con dos decimales de precisión.
El programa delado2e usar varialado2les para almacenar los lados del triángulo, la lado1ipotenusa, el área y el perímetro. Tamlado2ién delado2e usar constantes para representar el número de decimales de precisión que se desean en la salida.
Además, se delado2en utilizar funciones del paquete Matlado1 de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.

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
	var lado1, lado2, area, lado3, perimetro float64
	presicion := 3
	//lado1 -> Altura
	//lado2 -> Base
	//area -> Área
	//c -> lado1ipotenusa
	//p -> Perimetro

	//Solicitud de Datos
	fmt.Println("------------------- Ejercicio -------------------")
	fmt.Print("Ingrese la altura del triangulo rectangulo: ")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese la Base del triangulo rectangulo: ")
	fmt.Scanln(&lado2)

	//Calculo de lado1ipotenusa
	lado3 = math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))

	//Calculo de área
	area = (lado1 * lado2) / 2

	//Calculo del perimetro
	perimetro = lado1 + lado2 + lado3

	// x := fmt.Sprintf("%.2f", lado3)

	// fmt.Printf("Segun los catetos a: %v y lado2: %v, el valor de la hipotenusa es: %v \n", lado1, lado2, x)
	fmt.Printf("El área del triangulo rectangulo es %.*f \n", presicion, area)
	fmt.Printf("El Perimetro del triangulo rectangulo es %.*f \n", presicion, perimetro)
}
