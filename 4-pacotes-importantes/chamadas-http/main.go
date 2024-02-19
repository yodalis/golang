package main

import (
	"io"
	"net/http"
)

func main(){
	req, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()
	// Quando há o defer é como se pulasse esse comando 
	// e executasse ele depois de todas as outras linhas de código

	// Quando você quer que algo seja feito por último, você utiliza o defer

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))
	
	//Assim como os arquivos, é necessário fechar a conexão com o servidor
	// Para que o servidor saiba que a conexão foi fechada e evitar que fique aberta indefinidamente
}