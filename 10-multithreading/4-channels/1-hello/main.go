package main

import "fmt"

// Thread 1

func main() {
	canal := make(chan string) //canal vazio

	// Thread 2
	go func() {
		canal <- "Hello World" // preenchendo o canal
		// só é possível preencher dnv o canal quando esvaziarem ele
	}()

	// Thread 1
	msg := <-canal // esvaziando canal
	fmt.Println(msg)
}
