package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	server "github.com/yodalis/golang/6-banco-de-dados/desafio/server"
)

func main() {
	http.HandleFunc("/cotacao", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	currencyResp, err := server.FetchExchangeRate()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = server.InsertDataDB(db, currencyResp)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(currencyResp.USDBRL.Bid)
}
