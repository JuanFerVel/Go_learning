package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

func (l *ListaTareas) editarTarea(index int, tarea Tarea) {
	l.tareas[index] = tarea
}

func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	//Instlacia de Lista de Tareas
	lista := ListaTareas{}

	//Instancia de Bufui para ingresar tareas
	leer := bufio.NewReader(os.Stdin)

	//Listar todas las Tareas

	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar Tarea Completada\n",
			"3. Editar Tarea\n",
			"4. Eliminar Tarea\n",
			"5. Salir")
		fmt.Print("Digite la opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("------------------------------------------")
			fmt.Print("Ingrese el nombre de la Tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripción de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea Agregada Exitosamente")

		case 2:
			var index int
			fmt.Println("------------------------------------------")
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea Completada Exitosamente")

		case 3:
			var index int
			var t Tarea
			fmt.Println("------------------------------------------")
			fmt.Print("Ingrese el Indice de la tarea que desea Actualizar:")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la Tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripción de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea Editada Exitosamente")
		case 4:
			var index int
			fmt.Println("------------------------------------------")
			fmt.Print("Ingrese el Indice de la tarea que desea Actualizar:")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea Eliminada Exitosamente")
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("La opción ingresada no es correcta")
		}

		fmt.Println("==========================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s %s Completada: %t \n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("==========================================")

	}

}
