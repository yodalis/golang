package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	apiUrl     = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	dbTimeout  = 10 * time.Millisecond
	reqTimeout = 200 * time.Millisecond
)

type ExchangeRateResponse struct {
	USDBRL ExchangeRate `json:"USDBRL"`
}

type ExchangeRate struct {
	Bid string `json:"bid"`
}

func main() {
	http.HandleFunc("/cotacao", handler)
	http.ListenAndServe(":8080", nil)
}

func fetchExchangeRate() (*ExchangeRateResponse, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, reqTimeout)
	select {
	case <-ctx.Done():
		err := ctx.Err()
		if err == context.DeadlineExceeded {
			log.Fatal("Timeout server!")
		} else {
			log.Fatal("Error: ", err)
		}
	}
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var bodyJson ExchangeRateResponse
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		return nil, err
	}

	return &bodyJson, nil
}

func insertDataDB(db *sql.DB, currencyResp *ExchangeRateResponse) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	select {
	case <-ctx.Done():
		err := ctx.Err()
		if err == context.DeadlineExceeded {
			log.Fatal("Timeout server!")
		} else {
			log.Fatal("Error: ", err)
		}
	}
	defer cancel()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO currency_data (bid) VALUES (?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(currencyResp.USDBRL.Bid)
	if err != nil {
		return err
	}

	// select {
	// case <-time.After(dbTimeout):
	// 	fmt.Println("Operação demorou mais que 10ms")
	// case <-ctx.Done():
	// 	err := ctx.Err()
	// 	if err == context.DeadlineExceeded {
	// 		log.Fatal("Timeout server!")
	// 	} else {
	// 		log.Fatal("Error: ", err)
	// 	}
	// }

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	currencyResp, err := fetchExchangeRate()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = insertDataDB(db, currencyResp)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Println("final", currencyResp.USDBRL.Bid)

	json.NewEncoder(w).Encode(currencyResp.USDBRL.Bid)
}
