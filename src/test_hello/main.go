package main

import (
	"fmt"
	"log"

	"github.com/juanfervel/greetings"
)

func main() {
	log.SetPrefix("main(): ")

	message, err := greetings.Hello("Fer")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
