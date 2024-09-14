package main

import (
	"fmt"
	"net/http"

	"github.com/yodalis/golang/labs/googlecloud/webapp"
)

func main() {
	http.HandleFunc("/temperature", webapp.TemperatureHandler)
	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", nil)
}
