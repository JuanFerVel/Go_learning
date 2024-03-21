package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// canal := make(chan int)

	// canal <- 15
	// valor := <-canal

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com/",
	}

	ch := make(chan string)
	for _, api := range apis {
		go checkApi(api, ch)
	}
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// time.Sleep(5 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! Tomo %v segundos\n", elapsed.Seconds())
}

func checkApi(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		// fmt.Printf("ERROR: ¡%s Está caido! %v\n", api, err)
		ch <- fmt.Sprintf("ERROR: ¡%s Está caido! %v\n", api, err)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: ¡%s Está en funcionamiento!\n", api)
}
