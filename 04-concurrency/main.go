package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	start := time.Now()

	apis := []string{
		"https://dev.azure.com",
		"https://api.github.com",
		"https://api.somewhere.org",
		"https://outlook.office.com",
	}

	ch := make(chan string)

	for _, api := range apis {
		// go permite ejecutar la funcion en un hilo separado
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed.Seconds(), "seconds")
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Error: %s está caida", api)
		return
	}
	ch <- fmt.Sprintf("%s is up and running:", api)
}
