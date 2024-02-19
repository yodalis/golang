// Loopings
package main

func main() {
	// For basico
	for i := 0; i < 20; i++ {
		println(i)
	}

	// For range
	numbers := []string{"um", "dois", "tres "}
	for key, v := range numbers {
		println(key, v)
	}

	// loop infinito
	// for {
	// }
}
