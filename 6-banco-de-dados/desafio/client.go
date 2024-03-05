package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	clientTimeout = 300 * time.Millisecond
	apiUrl        = "http://localhost:8080/cotacao"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Println("Timeout client")
		return
	default:
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("cotacao.txt")
	fmt.Println("Criando arquivo...")
	if err != nil {
		panic(err)
	}

	valueString := string(body)
	valueString = strings.TrimSpace(valueString)
	valueString = strings.Trim(valueString, "\"")

	contentFile := fmt.Sprintf("Dólar: {%v}", valueString)

	_, err = file.WriteString(contentFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo criado com sucesso!")
	file.Close()
}
