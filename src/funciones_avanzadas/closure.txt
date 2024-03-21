package main

import "fmt"

func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := incrementer()

	fmt.Println(nextInt())
}
