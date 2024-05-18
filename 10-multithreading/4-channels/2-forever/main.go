package main

// Esse esquema "forever" segura processos até onde vc desejar
func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever
}
