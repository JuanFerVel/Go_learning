package animal

import "fmt"

type Animal interface {
	Sonido()
}
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace guau guau")
}

type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " hace miau")
}

func HacerSonido(animal Animal) {
	animal.Sonido()
}
